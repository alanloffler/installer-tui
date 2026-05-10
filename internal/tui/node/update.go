package node

import (
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/domain"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "ctrl-c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.Packages) {
			m.cursor++
		}
	case "space":
		if m.cursor == len(m.Packages) {
			return m, nil
		}
		if _, ok := m.selected[m.cursor]; ok {
			delete(m.selected, m.cursor)
		} else {
			m.selected[m.cursor] = struct{}{}
		}
	case "enter":
		if m.cursor == len(m.Packages) {
			return m, func() tea.Msg { return HomeMsg{} }
		}
		if len(m.selected) == 0 {
			return m, nil
		}

		out := make([]domain.Package, 0, len(m.selected))
		for i := range m.selected {
			out = append(out, m.Packages[i])
		}

		return m, func() tea.Msg { return DoneMsg{Selected: out} }
	}

	return m, nil
}
