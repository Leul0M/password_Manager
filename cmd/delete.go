package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"passwordmanager/db"
)

type deleteModel struct {
	cursor int
	creds  []db.Credential
}

func NewDeleteModel() tea.Model {
	creds, _ := db.GetAllCredentials()
	return deleteModel{
		creds: creds,
	}
}

func (m deleteModel) Init() tea.Cmd {
	return nil
}

func (m deleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "up":
					if m.cursor > 0 {
						m.cursor--
					}
				case "down":
					if m.cursor < len(m.creds)-1 {
						m.cursor++
					}
				case "enter":
					return confirmDeleteModel{cred: m.creds[m.cursor]}, nil
				case "esc":
					return NewMainMenuModel(), nil
			}
	}
	return m, nil
}

func (m deleteModel) View() string {
	s := HeaderStyle.Render("\n🗑️ Select a credential to delete:\n\n")

	for i, cred := range m.creds {
		cursor := " "
		if i == m.cursor {
			cursor = "➤"
		}
		line := fmt.Sprintf("%s %s | %s\n", cursor, cred.Email, cred.Note)
		if i == m.cursor {
			line = SelectedStyle.Render(line)
		} else {
			line = MenuItemStyle.Render(line)
		}
		s += line
	}

	s += PromptStyle.Render("\nPress Enter to select, Esc to cancel.\n")
	return s
}


type confirmDeleteModel struct {
	cred db.Credential
}

func (m confirmDeleteModel) Init() tea.Cmd {
	return nil
}

func (m confirmDeleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "y":
					db.DeleteCredential(m.cred.ID)
					return NewMainMenuModel(), nil
				case "n", "esc":
					return NewMainMenuModel(), nil
			}
	}
	return m, nil
}

func (m confirmDeleteModel) View() string {
	return PromptStyle.Render(fmt.Sprintf(
		"\n⚠️ Are you sure you want to delete:\nEmail: %s | Note: %s\n\nPress 'y' to confirm, 'n' or ESC to cancel.",
		m.cred.Email, m.cred.Note))
}

