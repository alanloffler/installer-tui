package done

import (
	"errors"

	"github.com/alanloffler/bubbletea/internal/installer"
	"github.com/alanloffler/bubbletea/internal/tui/styles"

	tea "charm.land/bubbletea/v2"
)

func (m Model) SectionSubtitle() string {
	return "Resúmen"
}

func (m Model) View() tea.View {
	s := ""

	for _, r := range m.Results {
		switch {
		case errors.Is(r.Err, installer.ErrAlreadyExists):
			s += styles.Fail.Render("⚠ "+r.Project.Name) + "  " + styles.Dim.Render(r.Err.Error()) + "\n\n"
		case r.Err != nil:
			s += styles.Fail.Render("⚠ "+r.Project.Name) + "  " + styles.Dim.Render(r.Err.Error()) + "\n\n"
		default:
			s += styles.OK.Render("✓ "+r.Project.Name) + "\n"
			if r.Project.UsageHint != "" {
				s += "  " + styles.Dim.Render(r.Project.UsageHint) + "\n\n"
			}
		}
	}

	s += "\n" + styles.Dim.Render("esc: inicio • q: salir")

	return tea.NewView(s)
}
