package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kkrull/bubbletea-hello/greet"
)

func main() {
	p := tea.NewProgram(greet.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Bubble Tea program failed: %v", err)
		os.Exit(1)
	}
}
