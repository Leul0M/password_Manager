package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"passwordmanager/db"
)

type listModel struct {
	data string
}

func NewListModel() tea.Model {
	data, _ := db.ListCredentials()
	return listModel{data: data}
}

func (m listModel) Init() tea.Cmd {
	return nil
}

func (m listModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "esc" {
				return NewMainMenuModel(), nil
			}
	}
	return m, nil
}

func (m listModel) View() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("79"))

	return style.Render("\nStored Credentials:\n" + m.data + "\nPress ESC to return to menu.")
}
