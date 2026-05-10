package done

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alanloffler/bubbletea/internal/installer"
	"github.com/alanloffler/bubbletea/internal/tui/styles"

	tea "charm.land/bubbletea/v2"
)

func (m Model) SectionSubtitle() string {
	return "Detalles de la instalación"
}

func (m Model) View() tea.View {
	s := ""

	for _, r := range m.Results {
		switch {
		case errors.Is(r.Err, installer.ErrAlreadyExists):
			s += styles.ErrorStyle.Render("✗ "+r.Name) + "  " + styles.MutedStyle.Render(r.Err.Error()) + "\n"
		case errors.Is(r.Err, installer.ErrPackageNotFound):
			s += styles.ErrorStyle.Render("✗ "+r.Name) + "  " + styles.MutedStyle.Render("paquete no encontrado") + "\n"
		case errors.Is(r.Err, installer.ErrNoProject):
			s += styles.ErrorStyle.Render("✗ "+r.Name) + "  " + styles.MutedStyle.Render(r.Err.Error()) + "\n"
		case r.Err != nil:
			s += styles.ErrorStyle.Render("✗ "+r.Name) + "  " + styles.MutedStyle.Render(r.Err.Error()) + "\n"
		default:
			s += styles.SuccessStyle.Render("✓ " + r.Name)
			if r.Version != "" {
				s += " " + styles.MutedStyle.Render(strings.ToLower(r.Name)+"@"+r.Version) + "\n"
			} else {
				s += "\n"
			}
			if r.UsageHint != "" {
				s += "  " + styles.MutedStyle.Render(r.UsageHint) + "\n"
			}
		}
	}

	s += "\n"

	for i, item := range m.items {
		cursor := "  "

		if m.cursor == i {
			cursor = styles.SelectedStyle.Render("▸ ")
			s += fmt.Sprintf("%s%s\n", cursor, styles.SelectedStyle.Render(item.Label))
		} else {
			s += fmt.Sprintf("%s%s\n", cursor, item.Label)
		}
	}

	s += "\n" + styles.HelpStyle.Render("j/k: navegar • enter: seleccionar • q: salir")

	return tea.NewView(s)
}
