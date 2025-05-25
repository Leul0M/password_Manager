package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	p := tea.NewProgram(NewMainMenuModel())
	if err := p.Start(); err != nil {
		panic(err)
	}
}
