package node

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/tui/styles"
)

func (m Model) SectionSubtitle() string {
	return "Elegí los paquetes que querés instalar"
}

func (m Model) View() tea.View {
	s := ""

	for i, p := range m.Packages {
		checkbox := styles.UnselectedStyle.Render("[ ]")
		name := styles.UnselectedStyle.Render(p.Name)

		cursor := " "
		if m.cursor == i {
			cursor = styles.UnselectedStyle.Render("▸")
		}

		_, isSelected := m.selected[i]
		if isSelected {
			checkbox = styles.SuccessStyle.Render("[x]")
			name = styles.SuccessStyle.Render(p.Name)
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, checkbox, name)
	}

	homeCursor := "  "
	if m.cursor == len(m.Packages) {
		homeCursor = styles.SelectedStyle.Render("▸ ")
		s += "\n" + homeCursor + styles.SelectedStyle.Render("Volver")
	} else {
		s += "\n" + homeCursor + styles.UnselectedStyle.Render("Volver")
	}

	s += "\n\n" + styles.HelpStyle.Render("j/k: navegar • space: seleccionar • enter: instalar • q: salir")

	return tea.NewView(s)
}
