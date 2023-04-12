package styles

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle   = lipgloss.NewStyle().Margin(1, 0).Foreground(lipgloss.Color("#aa6f5a"))
	VersionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff9671"))
	AuthorStyle  = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#ff9671"))
	HelpStyle    = lipgloss.NewStyle().Italic(true).Faint(true).Foreground(lipgloss.Color("#ffe6d6"))
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Italic(true)
	ReplStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("76")).Italic(true)
	PromptStyle  = ""
)
