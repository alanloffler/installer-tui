package selector

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/tui/styles"
)

func (m Model) SectionSubtitle() string {
	return "Elegí los proyectos que querés instalar"
}

func (m Model) View() tea.View {
	s := ""

	for i, p := range m.Projects {
		cursor := " "
		if m.cursor == i {
			cursor = "▸"
		}

		checkbox := "[ ]"
		name := p.Name

		_, isSelected := m.selected[i]
		if isSelected {
			checkbox = styles.Selected.Render("[x]")
			name = styles.Selected.Render(p.Name)
		}

		s += fmt.Sprintf("%s %s %s %s\n", cursor, checkbox, name, styles.Dim.Render(p.Repo))
	}

	s += "\n" + styles.Dim.Render("j/k: navegar • space: seleccionar • enter: instalar • q: salir")

	return tea.NewView(s)
}
