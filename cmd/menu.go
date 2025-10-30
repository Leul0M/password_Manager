package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menuModel struct {
	choices []string
	cursor  int
	width   int
	height  int
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
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
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
	header := HeaderStyle.Render("🗄️ Password Manager") + "\n" +
		SubtitleStyle.Render("Manage credentials securely")

	body := HelpTextStyle.Render("\nUse ↑/↓ or j/k, then Enter:\n\n")

	// Compute uniform width based on *menu items only*
	maxWidth := 0
	for _, choice := range m.choices {
		if w := lipgloss.Width(choice); w > maxWidth {
			maxWidth = w
		}
	}

	// Create consistent styles for items with the same width and centered text
	itemStyle := MenuItemStyle.Width(maxWidth).Align(lipgloss.Center)
	selStyle := SelectedStyle.Width(maxWidth).Align(lipgloss.Center).Inline(false)

	// Render each menu item consistently
	for i, choice := range m.choices {
		if i == m.cursor {
			body += selStyle.Render(choice) + "\n"
		} else {
			body += itemStyle.Render(choice) + "\n"
		}
	}

	footer := "\n" + HelpTextStyle.Render("Press ") +
		PromptStyle.Render("q") + HelpTextStyle.Render(" to quit") + "\n"

	content := lipgloss.JoinVertical(lipgloss.Center, header, body, footer)

	// Center entire block
	if m.width > 0 && m.height > 0 {
		return lipgloss.Place(
			m.width, m.height,
			lipgloss.Center, lipgloss.Center,
			ContainerStyle.Render(content),
		)
	}

	return ContainerStyle.Render(content)
}
