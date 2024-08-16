package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(InitialView())
	if _, err := p.Run(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(2)
	}
}
