package installing

import (
	"errors"

	"github.com/alanloffler/bubbletea/internal/installer"
	"github.com/alanloffler/bubbletea/internal/tui/styles"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Subtitle() string {
	return "Instalando los paquetes seleccionados"
}

func (m Model) View() tea.View {
	s := m.progress.View() + "\n\n"

	for i, p := range m.Queue {
		switch {
		case i < len(m.Done):
			r := m.Done[i]
			switch {
			case errors.Is(r.Err, installer.ErrAlreadyExists):
				s += styles.Warn.Render("⚠ "+p.Name) + " " + styles.Dim.Render(r.Err.Error()) + "\n"
			case r.Err != nil:
				s += styles.Fail.Render("⚠ "+p.Name) + " " + styles.Dim.Render(r.Err.Error()) + "\n"
			default:
				s += styles.OK.Render("✓ "+p.Name) + "\n"
			}
		case i == m.current:
			s += "• " + p.Name + styles.Dim.Render(" instalando...") + "\n"
		default:
			s += styles.Dim.Render("  "+p.Name+" pendiente") + "\n"
		}
	}

	return tea.NewView(s)
}
