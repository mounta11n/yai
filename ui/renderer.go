package ui

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const (
	exec_color    = "#FFFF00"
	config_color  = "#00D7FF"
	chat_color    = "#55FF00"
	help_color    = "#AAAAFF"
	error_color   = "#FF0000"
	warning_color = "#FFAA00"
	success_color = "#55AA00"
)

type Renderer struct {
	contentRenderer *glamour.TermRenderer
	successRenderer lipgloss.Style
	warningRenderer lipgloss.Style
	errorRenderer   lipgloss.Style
	helpRenderer    lipgloss.Style
}

func NewRenderer(options ...glamour.TermRendererOption) *Renderer {
	contentRenderer, err := glamour.NewTermRenderer(options...)
	if err != nil {
		return nil
	}

	successRenderer := lipgloss.NewStyle().Foreground(lipgloss.Color(success_color))
	warningRenderer := lipgloss.NewStyle().Foreground(lipgloss.Color(warning_color))
	errorRenderer := lipgloss.NewStyle().Foreground(lipgloss.Color(error_color))
	helpRenderer := lipgloss.NewStyle().Foreground(lipgloss.Color(help_color))

	return &Renderer{
		contentRenderer: contentRenderer,
		successRenderer: successRenderer,
		warningRenderer: warningRenderer,
		errorRenderer:   errorRenderer,
		helpRenderer:    helpRenderer,
	}
}

func (r *Renderer) RenderContent(in string) string {
	out, _ := r.contentRenderer.Render(in)

	return out
}

func (r *Renderer) RenderSuccess(in string) string {
	return r.successRenderer.Render(in)
}

func (r *Renderer) RenderWarning(in string) string {
	return r.warningRenderer.Render(in)
}

func (r *Renderer) RenderError(in string) string {
	return r.errorRenderer.Render(in)
}

func (r *Renderer) RenderHelp(in string) string {
	return r.helpRenderer.Render(in)
}

func (r *Renderer) RenderConfigMessage() string {
	welcome := "Hallo!  \n\n"
	welcome += "Es konnte keine Konfigurationsdatei gefunden werden, bitte gib einen API Schlüssel ein"

	return welcome
}

func (r *Renderer) RenderHelpMessage() string {
	help := "**Tastenkombinationen**\n"
	help += "- `↑`/`↓` : In History navigieren\n"
	help += "- `Tab`   : Modus wechseln zwischen `Exec` & `Chat`\n"
	help += "- `Strg+h`: Hilfe anzeigen\n"
	help += "- `Strg+s`: Einstellungen\n"
	help += "- `Strg+r`: Terminal löschen und Diskussionsverlauf zurücksetzen\n"
	help += "- `Strg+l`: Terminal löschen, aber Diskussionsverlauf beibehalten\n"
	help += "- `Strg+c`: Beende oder unterbreche Befehl\n"

	return help
}
