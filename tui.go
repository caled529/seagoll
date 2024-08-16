package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

func InitialView() tea.Model {
	return seagollViewer{
		grid:   &Grid{tileMap: make(map[int]map[int]bool)},
		paused: true,
		styleMenu: gloss.NewStyle().
			Align(gloss.Center),
		styleView: gloss.NewStyle().
			Border(gloss.NormalBorder()),
	}
}

type seagollViewer struct {
	cursorX   int
	cursorY   int
	grid      *Grid
	paused    bool
	styleMenu gloss.Style
	styleView gloss.Style
	viewportX int
	viewportY int
}

func (sv seagollViewer) Init() tea.Cmd { return nil }

func (sv seagollViewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return sv, tea.Quit
		case "up", "k":
			if sv.cursorY < 1 {
				sv.viewportY -= 1
			} else {
				sv.cursorY--
			}
		case "down", "j":
			if sv.cursorY >= sv.styleView.GetHeight()-1 {
				sv.viewportY += 1
			} else {
				sv.cursorY++
			}
		case "left", "h":
			if sv.cursorX < 1 {
				sv.viewportX -= 1
			} else {
				sv.cursorX--
			}
		case "right", "l":
			if sv.cursorX >= sv.styleView.GetWidth()/2-1 {
				sv.viewportX += 1
			} else {
				sv.cursorX++
			}
		case "n":
			if sv.paused {
				sv.grid.NextGeneration()
			}
		case "enter":
			if sv.paused {
				sv.grid.ToggleAt(sv.viewportX+sv.cursorX, sv.viewportY+sv.cursorY)
			}
		case " ":
			sv.paused = !sv.paused
		}
	case tea.WindowSizeMsg:
		sv.styleMenu = sv.styleMenu.
			Width(msg.Width)
		msg.Width -= 2
		msg.Height -= 7
		sv.styleView = sv.styleView.
			Width(msg.Width).
			Height(msg.Height)
		if sv.cursorX >= msg.Width/2 {
			sv.viewportX += sv.cursorX - msg.Width/2 + 1
			sv.cursorX = msg.Width/2 - 1
		}
		if sv.cursorY >= msg.Height {
			sv.viewportY += sv.cursorY
			sv.cursorX = msg.Height - 1
		}
	}
	return sv, nil
}

func (sv seagollViewer) View() string {
	s := sv.styleMenu.Render("Conway's Game of Life")
	s += fmt.Sprintf("\nCoordinates: x=%d, y=%d", sv.viewportX+sv.cursorX, sv.viewportY+sv.cursorY)
	gridView := strings.Split(sv.grid.BoundedView(sv.viewportX, sv.viewportY,
		sv.styleView.GetWidth()/2, sv.styleView.GetHeight()), "\n")
	cursorText := "❱❰"
	if sv.grid.IsAliveAt(sv.viewportX+sv.cursorX, sv.viewportY+sv.cursorY) {
		cursorText = gloss.NewStyle().
			Background(gloss.Color("8")).
			Render(cursorText)
	}
	cursorLine := []rune(gridView[sv.cursorY])
	if len(cursorLine) >= 2 {
		gridView[sv.cursorY] = string(cursorLine[:sv.cursorX*2]) + cursorText + string(cursorLine[(sv.cursorX+1)*2:])
	}
	s += "\n" + sv.styleView.Render(strings.Join(gridView, "\n"))
	if sv.paused {
		s += sv.styleMenu.Render("\n(Press Spacebar to Run)    (Press Enter to toggle a tile)\n(Use ←↓↑→ (or hjkl) to move the cursor)  (Press n to advance)")
	} else {
		s += sv.styleMenu.Render("\n(Press Spacebar to Pause)\n(Use ←↓↑→ (or hjkl) to move the cursor)")
	}
	s += sv.styleMenu.Render("\n(Press q to quit at any time)")
	return s
}
