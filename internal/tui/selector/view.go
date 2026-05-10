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
		checkbox := styles.UnselectedStyle.Render("[ ]")
		name := styles.UnselectedStyle.Render(p.Name)
		description := styles.SubtextStyle.Render(p.Description)
		repo := styles.MutedStyle.PaddingLeft(6).Render(p.Repo)

		cursor := " "
		if m.cursor == i {
			cursor = styles.UnselectedStyle.Render("▸")
		}

		_, isSelected := m.selected[i]
		if isSelected {
			checkbox = styles.SuccessStyle.Render("[x]")
			name = styles.SuccessStyle.Render(p.Name)
		}

		s += fmt.Sprintf("%s %s %s %s\n", cursor, checkbox, name, description)
		s += fmt.Sprintf("%s\n", repo)
	}

	homeCursor := "  "

	if m.cursor == len(m.Projects) {
		homeCursor = styles.SelectedStyle.Render("▸ ")
		s += "\n" + homeCursor + styles.SelectedStyle.Render("Volver")
	} else {
		s += "\n" + homeCursor + styles.UnselectedStyle.Render("Volver")
	}

	s += "\n\n" + styles.HelpStyle.Render("j/k: navegar • space: seleccionar • enter: instalar • q: salir")

	return tea.NewView(s)
}
