package ai

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/mounta11n/yai/config"
	"github.com/mounta11n/yai/system"

	"github.com/sashabaranov/go-openai"
)

const noexec = "[noexec]"

type Engine struct {
	mode         EngineMode
	config       *config.Config
	client       *openai.Client
	execMessages []openai.ChatCompletionMessage
	chatMessages []openai.ChatCompletionMessage
	channel      chan EngineChatStreamOutput
	pipe         string
	running      bool
}

func NewEngine(mode EngineMode, config *config.Config) (*Engine, error) {
	clientConfig := openai.DefaultConfig(config.GetAiConfig().GetKey())

	if config.GetAiConfig().GetProxy() != "" {
		proxyUrl, err := url.Parse(config.GetAiConfig().GetProxy())
		if err != nil {
			return nil, err
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}

		clientConfig.HTTPClient = &http.Client{
			Transport: transport,
		}
	}

	if config.GetAiConfig().GetBaseUrl() != "" {
		clientConfig.BaseURL = config.GetAiConfig().GetBaseUrl()
	}

	client := openai.NewClientWithConfig(clientConfig)
	return &Engine{
		mode:         mode,
		config:       config,
		client:       client,
		execMessages: make([]openai.ChatCompletionMessage, 0),
		chatMessages: make([]openai.ChatCompletionMessage, 0),
		channel:      make(chan EngineChatStreamOutput),
		pipe:         "",
		running:      false,
	}, nil
}

func (e *Engine) SetMode(mode EngineMode) *Engine {
	e.mode = mode

	return e
}

func (e *Engine) GetMode() EngineMode {
	return e.mode
}

func (e *Engine) GetChannel() chan EngineChatStreamOutput {
	return e.channel
}

func (e *Engine) SetPipe(pipe string) *Engine {
	e.pipe = pipe

	return e
}

func (e *Engine) Interrupt() *Engine {
	e.channel <- EngineChatStreamOutput{
		content:    "[Interrupt]",
		last:       true,
		interrupt:  true,
		executable: false,
	}

	e.running = false

	return e
}

func (e *Engine) Clear() *Engine {
	if e.mode == ExecEngineMode {
		e.execMessages = []openai.ChatCompletionMessage{}
	} else {
		e.chatMessages = []openai.ChatCompletionMessage{}
	}

	return e
}

func (e *Engine) Reset() *Engine {
	e.execMessages = []openai.ChatCompletionMessage{}
	e.chatMessages = []openai.ChatCompletionMessage{}

	return e
}

func (e *Engine) ExecCompletion(input string) (*EngineExecOutput, error) {
	ctx := context.Background()

	e.running = true

	e.appendUserMessage(input)

	resp, err := e.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:     e.config.GetAiConfig().GetModel(),
			MaxTokens: e.config.GetAiConfig().GetMaxTokens(),
			Messages:  e.prepareCompletionMessages(),
		},
	)
	if err != nil {
		return nil, err
	}

	content := resp.Choices[0].Message.Content
	e.appendAssistantMessage(content)

	var output EngineExecOutput
	err = json.Unmarshal([]byte(content), &output)
	if err != nil {
		re := regexp.MustCompile(`\{.*?\}`)
		match := re.FindString(content)
		if match != "" {
			err = json.Unmarshal([]byte(match), &output)
			if err != nil {
				return nil, err
			}
		} else {
			output = EngineExecOutput{
				Command:     "",
				Explanation: content,
				Executable:  false,
			}
		}
	}

	return &output, nil
}

func (e *Engine) ChatStreamCompletion(input string) error {
	ctx := context.Background()

	e.running = true

	e.appendUserMessage(input)

	req := openai.ChatCompletionRequest{
		Model:     e.config.GetAiConfig().GetModel(),
		MaxTokens: e.config.GetAiConfig().GetMaxTokens(),
		Messages:  e.prepareCompletionMessages(),
		Stream:    true,
	}

	stream, err := e.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return err
	}
	defer stream.Close()

	var output string

	for {
		if e.running {
			resp, err := stream.Recv()

			if errors.Is(err, io.EOF) {
				executable := false
				if e.mode == ExecEngineMode {
					if !strings.HasPrefix(output, noexec) && !strings.Contains(output, "\n") {
						executable = true
					}
				}

				e.channel <- EngineChatStreamOutput{
					content:    "",
					last:       true,
					executable: executable,
				}
				e.running = false
				e.appendAssistantMessage(output)

				return nil
			}

			if err != nil {
				e.running = false
				return err
			}

			delta := resp.Choices[0].Delta.Content

			output += delta

			e.channel <- EngineChatStreamOutput{
				content: delta,
				last:    false,
			}

			// time.Sleep(time.Microsecond * 100)
		} else {
			stream.Close()

			return nil
		}
	}
}

func (e *Engine) appendUserMessage(content string) *Engine {
	if e.mode == ExecEngineMode {
		e.execMessages = append(e.execMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		})
	} else {
		e.chatMessages = append(e.chatMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		})
	}

	return e
}

func (e *Engine) appendAssistantMessage(content string) *Engine {
	if e.mode == ExecEngineMode {
		e.execMessages = append(e.execMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
	} else {
		e.chatMessages = append(e.chatMessages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
	}

	return e
}

func (e *Engine) prepareCompletionMessages() []openai.ChatCompletionMessage {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: e.prepareSystemPrompt(),
		},
	}

	if e.pipe != "" {
		messages = append(
			messages,
			openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: e.preparePipePrompt(),
			},
		)
	}

	if e.mode == ExecEngineMode {
		messages = append(messages, e.execMessages...)
	} else {
		messages = append(messages, e.chatMessages...)
	}

	return messages
}

func (e *Engine) preparePipePrompt() string {
	return fmt.Sprintf("Ich werde an den folgenden Eingaben arbeiten: %s", e.pipe)
}

func (e *Engine) prepareSystemPrompt() string {
	var bodyPart string
	if e.mode == ExecEngineMode {
		bodyPart = e.prepareSystemPromptExecPart()
	} else {
		bodyPart = e.prepareSystemPromptChatPart()
	}

	return fmt.Sprintf("%s\n%s", bodyPart, e.prepareSystemPromptContextPart())
}

func (e *Engine) prepareSystemPromptExecPart() string {
	return "Du bist Yai, ein mächtiger Terminal-Assistent, der eine JSON-Datei mit einer Befehlszeile für meine Eingabe generiert.\n" +
	"Du wirst immer mit der folgenden json-Struktur antworten: {\"cmd\":\"der Befehl\", \"exp\": \"eine Erklärung\", \"exec\": true}.\n" +
		"Deine Antwort wird immer nur die json-Struktur enthalten, niemals Ratschläge oder zusätzliche Details oder Informationen hinzufügen, auch wenn ich die gleiche Frage schon einmal gestellt habe.\n" +
		"Das Feld cmd enthält einen einzeiligen Befehl (verwende keine neuen Zeilen, sondern Trennzeichen wie && und ;\n" +
		"Das Feld exp enthält eine kurze Erklärung des Befehls, wenn es dir gelungen ist, einen ausführbaren Befehl zu generieren, andernfalls enthält es den Grund für deinen Fehler.\n" +
		"Das Feld exec enthält true, wenn es dir gelungen ist, einen ausführbaren Befehl zu erzeugen, sonst false." +
		"\n" +
		"Examples:\n" +
		"Me: alle Dateien in meinem Homeverzeichnis auflisten\n" +
		"Yai: {\"cmd\":\"ls ~\", \"exp\": \"Auflistung aller Dateien in deinem Home-Verzeichnis\", \"exec\\: true}\n" +
		"Me: Bitte liste alle Pods aller Namespaces auf\n" +
		"Yai: {\"cmd\":\"kubectl get pods --all-namespaces\", \"exp\": \"Pods bilden alle k8s Namespaces\", \"exec\": true}\n" +
		"Me: Wie geht es dir?\n" +
		"Yai: {\"cmd\":\"\", \"exp\": \"Vielen Dank, aber ich kann keinen Befehl dafür erstellen. Benutze den Chat-Modus, wenn du mit mir diskutieren willst.\", \"exec\": false}"
}

func (e *Engine) prepareSystemPromptChatPart() string {
	return "Du bist Yai, ein mächtiger Terminal-Assistent, der von Yazan erschaffen wurde.\n" +
		"Du wirst auf die hilfreichste Art und Weise antworten.\n" +
		"Formatiere deine Antwort immer im Markdown-Format.\n\n" +
		"Beispiel:\n" +
		"Me: Wie viel ist 2+2 ?\n" +
		"Yai: Die Antwort für `2+2` ist `4`\n" +
		"Me: und nochmal +2?\n" +
		"Yai: Die Summe wäre dann `6`\n"
}

func (e *Engine) prepareSystemPromptContextPart() string {
	part := "Mein Kontext: "

	if e.config.GetSystemConfig().GetOperatingSystem() != system.UnknownOperatingSystem {
		part += fmt.Sprintf("mein Betriebssystem ist %s, ", e.config.GetSystemConfig().GetOperatingSystem().String())
	}
	if e.config.GetSystemConfig().GetDistribution() != "" {
		part += fmt.Sprintf("meine Disro ist %s, ", e.config.GetSystemConfig().GetDistribution())
	}
	if e.config.GetSystemConfig().GetHomeDirectory() != "" {
		part += fmt.Sprintf("mein Homeverzeichnis ist %s, ", e.config.GetSystemConfig().GetHomeDirectory())
	}
	if e.config.GetSystemConfig().GetShell() != "" {
		part += fmt.Sprintf("meine Shell ist %s, ", e.config.GetSystemConfig().GetShell())
	}
	if e.config.GetSystemConfig().GetShell() != "" {
		part += fmt.Sprintf("mein Editor ist %s, ", e.config.GetSystemConfig().GetEditor())
	}
	part += "berücksichtige dies. "

	if e.config.GetUserConfig().GetPreferences() != "" {
		part += fmt.Sprintf("Außerdem noch, %s.", e.config.GetUserConfig().GetPreferences())
	}

	return part
}
