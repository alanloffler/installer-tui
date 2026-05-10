package header

import (
	"strings"

	"github.com/alanloffler/bubbletea/internal/tui/styles"
)

var (
	titleStyle    = styles.HeadingStyle.Padding(0, 2)
	versionStyle  = styles.SubtextStyle
	subtitleStyle = styles.SubtextStyle.Bold(true).Italic(true).PaddingLeft(2)
)

func Render(appTitle, appSubtitle, sectionSubtitle string) string {
	var b strings.Builder
	b.WriteString(titleStyle.Render(appTitle))
	b.WriteString(versionStyle.Render(appSubtitle))
	b.WriteString("\n\n")

	if sectionSubtitle != "" {
		b.WriteString(subtitleStyle.Render(sectionSubtitle))
		b.WriteString("\n")
	}

	b.WriteString("\n")

	return b.String()
}
