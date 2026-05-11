package node

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/installer-tui/internal/tui/browser"
	"github.com/alanloffler/installer-tui/internal/tui/styles"
)

const cursorSlotsPerPkg = 2

func (m Model) SectionSubtitle() string {
	return "Elegí los paquetes que querés instalar"
}

func (m Model) View() tea.View {
	s := ""

	for i, p := range m.Packages {
		mainPos := cursorSlotsPerPkg * i
		repoPos := cursorSlotsPerPkg*i + 1

		checkbox := styles.UnselectedStyle.Render("[ ]")
		name := styles.UnselectedStyle.Render(p.Name)
		description := styles.SubtextStyle.Render(p.Description)
		repo := styles.MutedStyle.PaddingLeft(6).Render(p.Repo)
		if p.Repo == "" {
			repo = styles.MutedStyle.PaddingLeft(6).Render("✗ no disponible")
		}

		cursor := " "
		if m.cursor == mainPos {
			cursor = styles.UnselectedStyle.Render("▸")
		}

		if m.cursor == repoPos {
			repo = styles.UnselectedStyle.Render("▸")
			if browser.IsLink(p.Repo) {
				repo += styles.LinkStyle.PaddingLeft(5).Render(p.Repo)
			} else {
				repo += styles.MutedStyle.PaddingLeft(5).Render("✗ no disponible")
			}
		}

		_, isSelected := m.selected[i]
		if isSelected {
			checkbox = styles.SuccessStyle.Render("[x]")
			name = styles.SuccessStyle.Render(p.Name)
		}

		s += fmt.Sprintf("%s %s %s %s\n", cursor, checkbox, name, description)
		s += fmt.Sprintf("%s\n", repo)
	}

	s += "\n"
	for i, item := range m.items {
		c := "  "
		pos := cursorSlotsPerPkg*len(m.Packages) + i
		if m.cursor == pos {
			c = styles.SelectedStyle.Render("▸ ")
			s += c + styles.SelectedStyle.Render(item.Label)
		} else {
			s += c + styles.UnselectedStyle.Render(item.Label)
		}
	}

	s += "\n\n" + styles.HelpStyle.Render("j/k: navegar • space: seleccionar • enter: instalar • q: salir")

	return tea.NewView(s)
}
