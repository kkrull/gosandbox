package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	fmt.Printf("Hello World!\n")
}

type model struct {
	Message string
}

func initialModel() *model {
	return &model{Message: "Hello World!"}
}

func (m model) Init() tea.Cmd {
	return nil
}
