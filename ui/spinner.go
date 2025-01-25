package ui

import (
	"fmt"
	"math/rand"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

var loadingMessages = []string{
	"Lass Mich Kurz Überlegen ...",
	"Kleinen Moment ...",
	"Am Nachdenken ...",
	"Am Laden ...",
	"Neuronen Feuern ...",
	"Synapsen Arbeiten ...",
	"Sekunde Bitte ...",
	"Am Berechnen ...",
	"Am Verarbeiten ...",
	"Bitte Warten ...",
	"Neuronen Feuern ...",
	"Synapsen Arbeiten ...",
}

type Spinner struct {
	message string
	spinner spinner.Model
}

func NewSpinner() *Spinner {
	spin := spinner.New()
	spin.Spinner = spinner.MiniDot

	return &Spinner{
		message: loadingMessages[rand.Intn(len(loadingMessages))],
		spinner: spin,
	}
}

func (s *Spinner) Update(msg tea.Msg) (*Spinner, tea.Cmd) {
	var updateCmd tea.Cmd
	s.spinner, updateCmd = s.spinner.Update(msg)

	return s, updateCmd
}

func (s *Spinner) View() string {
	return fmt.Sprintf(
		"\n  %s %s...",
		s.spinner.View(),
		s.spinner.Style.Render(s.message),
	)
}

func (s *Spinner) Tick() tea.Msg {
	return s.spinner.Tick()
}
