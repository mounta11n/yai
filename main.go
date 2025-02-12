package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/mounta11n/yai/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	input, err := ui.NewUIInput()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tea.NewProgram(ui.NewUi(input)).Run(); err != nil {
		log.Fatal(err)
	}
}
