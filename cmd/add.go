package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"passwordmanager/db"
)

type addModel struct {
	step     int
	email    string
	password string
	note     string
	input    string
}

func NewAddModel() tea.Model {
	return addModel{}
}

func (m addModel) Init() tea.Cmd {
	return nil
}

func (m addModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "enter":
					switch m.step {
						case 0:
							m.email = m.input
							m.input = ""
							m.step++
						case 1:
							m.password = m.input
							m.input = ""
							m.step++
						case 2:
							m.note = m.input
							db.AddCredential(m.email, m.password, m.note)
							return NewMainMenuModel(), nil
					}
						case "esc":
							return NewMainMenuModel(), nil
						default:
							m.input += msg.String()
			}
	}
	return m, nil
}

func (m addModel) View() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("212"))

	var prompt string
	switch m.step {
		case 0:
			prompt = "Enter Email: "
		case 1:
			prompt = "Enter Password: "
		case 2:
			prompt = "Enter Note: "
	}

	return style.Render(fmt.Sprintf("%s%s", prompt, m.input))
}
