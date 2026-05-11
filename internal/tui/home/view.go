package home

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/tui/styles"
)

func (Model) SectionSubtitle() string {
	return "Seleccioná el tipo de instalación"
}

func (m Model) View() tea.View {
	s := ""

	for i, item := range m.items {
		cursor := " "

		if m.cursor == i {
			cursor = "▸"
			s += fmt.Sprintf("%s %s\n", styles.SelectedStyle.Render(cursor), styles.SelectedStyle.Render(item.Label))
		} else {
			s += fmt.Sprintf("%s %s\n", cursor, item.Label)
		}
	}

	s += "\n" + m.help.View(keys)

	return tea.NewView(s)
}
