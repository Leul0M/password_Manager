package cmd

import "github.com/charmbracelet/lipgloss"

var (
    // Modern, cohesive palette (inspired by Catppuccin)
    PrimaryColor      = lipgloss.Color("#89B4FA") // Blue
    AccentColor       = lipgloss.Color("#F38BA8") // Rose
    SuccessColor      = lipgloss.Color("#A6E3A1") // Green
    WarningColor      = lipgloss.Color("#F9E2AF") // Yellow
    MutedColor        = lipgloss.Color("#9399B2") // Muted gray
    SurfaceColor      = lipgloss.Color("#1E1E2E") // Base surface
    SurfaceAltColor   = lipgloss.Color("#181825") // Darker surface
    HighlightColor    = lipgloss.Color("#B4BEFE") // Lavender highlight

    // Header/title
    HeaderStyle       = lipgloss.NewStyle().Foreground(AccentColor).Bold(true)
    SubtitleStyle     = lipgloss.NewStyle().Foreground(MutedColor).Italic(true)

    // Container/card
    ContainerStyle    = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(PrimaryColor).
        Padding(1, 2).
        Margin(1, 2)

    // Menu items
    MenuItemStyle     = lipgloss.NewStyle().Foreground(MutedColor)
    SelectedStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("#11111B")).Background(PrimaryColor).Bold(true)
    CursorStyle       = lipgloss.NewStyle().Foreground(PrimaryColor).Bold(true)

    // Prompts and status
    PromptStyle       = lipgloss.NewStyle().Foreground(AccentColor).Bold(true)
    HelpTextStyle     = lipgloss.NewStyle().Foreground(MutedColor)
    SuccessStyle      = lipgloss.NewStyle().Foreground(SuccessColor).Bold(true)
    WarningTextStyle  = lipgloss.NewStyle().Foreground(WarningColor).Bold(true)
)
