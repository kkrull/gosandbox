package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(newGreeterModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Bubble Tea program failed: %v", err)
		os.Exit(1)
	}
}

type greeterModel struct {
	Message string
}

func newGreeterModel() *greeterModel {
	return &greeterModel{Message: "Hello World!"}
}

func (m greeterModel) Init() tea.Cmd {
	return nil
}

func (m greeterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m greeterModel) View() string {
	return fmt.Sprintf("%s\nPress q or Ctrl+C to quit.", m.Message)
}
