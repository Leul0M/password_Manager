package cmd

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor    = lipgloss.Color("#00C897") // Teal
	AccentColor     = lipgloss.Color("#FF5733") // Orange
	WarningColor    = lipgloss.Color("#FF3860") // Red
	HighlightColor  = lipgloss.Color("#FFD700") // Gold

	HeaderStyle     = lipgloss.NewStyle().Foreground(AccentColor).Bold(true).Underline(true)
	MenuItemStyle   = lipgloss.NewStyle().Foreground(PrimaryColor)
	SelectedStyle   = lipgloss.NewStyle().Foreground(HighlightColor).Bold(true)
	PromptStyle     = lipgloss.NewStyle().Foreground(WarningColor).Bold(true)
	SuccessStyle    = lipgloss.NewStyle().Foreground(PrimaryColor).Bold(true)
)
