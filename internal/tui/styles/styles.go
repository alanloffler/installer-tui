package styles

import "github.com/charmbracelet/lipgloss"

var (
	Title    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).Background(lipgloss.Color("#7C3AED")).Bold(true).Padding(0, 1)
	Dim      = lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280"))
	OK       = lipgloss.NewStyle().Foreground(lipgloss.Color("#22C55E"))
	Fail     = lipgloss.NewStyle().Foreground(lipgloss.Color("#EF4444"))
	Warn     = lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B"))
	Selected = lipgloss.NewStyle().Foreground(lipgloss.Color("#A78BFA")).Bold(true)
)
