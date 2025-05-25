package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type menuModel struct {
	choices []string
	cursor  int
}

func NewMainMenuModel() tea.Model {
	return menuModel{
		choices: []string{"Add Credential", "List Credentials", "Delete Credential", "Exit"},
		cursor:  0,
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}
func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "up", "k":
					if m.cursor > 0 {
						m.cursor--
					}
				case "down", "j":
					if m.cursor < len(m.choices)-1 {
						m.cursor++
					}
				case "enter":
					switch m.cursor {
						case 0:
							return NewAddModel(), nil
						case 1:
							return NewListModel(), nil
						case 2:
							return NewDeleteModel(), nil
						case 3:
							// Exit the program properly
							return m, tea.Quit
					}
						case "q", "esc":
							// Exit the program properly
							return m, tea.Quit
			}
	}
	return m, nil
}


func (m menuModel) View() string {
	s := HeaderStyle.Render("\n🗄️ Password Manager\n")
	s += MenuItemStyle.Render("\nUse arrow keys ↑ ↓ and press Enter:\n\n")

	for i, choice := range m.choices {
		cursor := " "
		if i == m.cursor {
			cursor = "➤"
		}
		line := fmt.Sprintf("%s %s\n", cursor, choice)
		if i == m.cursor {
			line = SelectedStyle.Render(line)
		} else {
			line = MenuItemStyle.Render(line)
		}
		s += line
	}

	s += PromptStyle.Render("\nPress q to quit.\n")
	return s
}
