package greet

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func NewModel() *greeterModel {
	return &greeterModel{Message: "Hello World!"}
}

/* bubbletea.Model */

type greeterModel struct {
	Message string
}

func (model greeterModel) Init() tea.Cmd {
	return nil
}

func (model greeterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return model, tea.Quit
		}
	}

	return model, nil
}

func (greeter greeterModel) View() string {
	return fmt.Sprintf("%s\nPress q or Ctrl+C to quit.", greeter.Message)
}
