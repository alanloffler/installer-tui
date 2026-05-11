package header

import (
	"os"
	"runtime"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/alanloffler/bubbletea/internal/tui/styles"
)

var (
	titleStyle = styles.HeadingStyle.Padding(0, 2, 0, 2).Border(lipgloss.Border{
		TopLeft: "┌", Top: "─", TopRight: "┐",
		Right:      "│",
		BottomLeft: "└", Bottom: "─", BottomRight: "┘",
		Left: "│"},
	).BorderForeground(styles.ColorMauve)
	subtitleStyle = styles.SubtextStyle.Bold(true).Italic(true)
)

func Render(appTitle, appSubtitle, sectionSubtitle string) string {
	var b strings.Builder
	b.WriteString(titleStyle.Render(appTitle))
	b.WriteString("\n")

	b.WriteString(styles.PrimaryStyle.Render("Version: "))
	b.WriteString(styles.SubtextStyle.Render(appSubtitle) + "\n")
	b.WriteString(styles.PrimaryStyle.Render("OS: "))
	b.WriteString(styles.SubtextStyle.Render(runtime.GOOS) + "\n")

	dir, _ := os.Getwd()
	b.WriteString(styles.PrimaryStyle.Render("Path: "))
	b.WriteString(styles.SubtextStyle.Render(dir) + "\n\n")

	if sectionSubtitle != "" {
		b.WriteString(subtitleStyle.Render(sectionSubtitle))
		b.WriteString("\n")
	}

	b.WriteString("\n")

	return b.String()
}
