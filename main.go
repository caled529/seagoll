package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(InitialView(time.Second / 4))
	if _, err := p.Run(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(2)
	}
}
