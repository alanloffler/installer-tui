package header

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#55d286")).Bold(true).Padding(0, 2)
	subtitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9ca3af")).Italic(true).PaddingLeft(2)
)

func Render(appTitle, appSubtitle, sectionSubtitle string) string {
	var b strings.Builder
	b.WriteString(titleStyle.Render(appTitle))
	b.WriteString(subtitleStyle.Render(appSubtitle))
	b.WriteString("\n\n")

	if sectionSubtitle != "" {
		b.WriteString(subtitleStyle.Render(sectionSubtitle))
		b.WriteString("\n")
	}

	b.WriteString("\n")

	return b.String()
}
