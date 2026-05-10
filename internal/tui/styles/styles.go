package styles

import "github.com/charmbracelet/lipgloss"

var (
	ColorGreen    = lipgloss.Color("#9ccfd8")
	ColorLavender = lipgloss.Color("#c4a7e7")
	ColorMauve    = lipgloss.Color("#ebbcba")
	ColorMauve2   = lipgloss.Color("#675d70")
	// ColorOverlay  = lipgloss.Color("#6e6a86") dark zinc
	// ColorPeach    = lipgloss.Color("#f6c177") light orange
	ColorRed     = lipgloss.Color("#eb6f92")
	ColorSubtext = lipgloss.Color("#908caa")
	ColorText    = lipgloss.Color("#e0def4")
	ColorYellow  = lipgloss.Color("#f1ca93")
)

var (
	HeadingStyle = lipgloss.NewStyle().Foreground(ColorMauve).Bold(true)
	SubtextStyle = lipgloss.NewStyle().Foreground(ColorSubtext)
	HelpStyle    = lipgloss.NewStyle().Foreground(ColorSubtext)

	ErrorStyle   = lipgloss.NewStyle().Foreground(ColorRed)
	SuccessStyle = lipgloss.NewStyle().Foreground(ColorGreen)
	WarningStyle = lipgloss.NewStyle().Foreground(ColorYellow)

	SelectedStyle   = lipgloss.NewStyle().Foreground(ColorLavender).Bold(true)
	UnselectedStyle = lipgloss.NewStyle().Foreground(ColorText)
	MutedStyle      = lipgloss.NewStyle().Foreground(ColorMauve2)
)
