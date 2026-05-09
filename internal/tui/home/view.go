package home

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/tui/styles"
)

func (Model) SectionSubtitle() string {
	return "Seleccioná tipo de instalación"
}

func (m Model) View() tea.View {
	s := ""

	for i, item := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = "▸"
		}

		s += fmt.Sprintf("%s %s\n", cursor, item.label)
	}

	s += "\n" + styles.Dim.Render("j/k: navegar • enter: seleccionar • q: salir")

	return tea.NewView(s)
}
