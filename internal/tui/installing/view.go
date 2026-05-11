package installing

import (
	"errors"

	"github.com/alanloffler/installer-tui/internal/installer"
	"github.com/alanloffler/installer-tui/internal/tui/styles"

	tea "charm.land/bubbletea/v2"
)

func (m Model) SectionSubtitle() string {
	return "Instalación"
}

func (m Model) View() tea.View {
	s := m.progress.View() + "\n\n"

	if m.showPM {
		if m.PM != "" {
			s += styles.HighlightStyle.Render(m.PM) + " " + styles.SubtextStyle.Render("detectado") + "\n\n"
		} else {
			s += styles.ErrorStyle.Render("Package manager no detectado") + " "
			s += styles.SubtextStyle.Render("Debes iniciar un proyecto") + "\n\n"
		}
	}

	for i, p := range m.Queue {
		switch {
		case i < len(m.Done):
			r := m.Done[i]
			switch {
			case errors.Is(r.Err, installer.ErrAlreadyExists):
				s += styles.WarningStyle.Render("⚠ "+p.Name) + "\n"
			case r.Err != nil:
				s += styles.ErrorStyle.Render("✗ "+p.Name) + "\n"
			default:
				s += styles.SuccessStyle.Render("✓ "+p.Name) + "\n"
			}
		case i == m.current:
			s += m.spinner.View() + " " + styles.UnselectedStyle.Render(p.Name) + "\n"
		default:
			s += styles.MutedStyle.Render("  "+p.Name) + "\n"
		}
	}

	if m.finished {
		s += "\n" + styles.SelectedStyle.Bold(false).Render("Presiona Enter para continuar")
	}

	return tea.NewView(s)
}
