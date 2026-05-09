package selector

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
		if m.cursor < len(m.Projects)-1 {
			m.cursor++
		}
	case "space":
		if _, ok := m.selected[m.cursor]; ok {
			delete(m.selected, m.cursor)
		} else {
			m.selected[m.cursor] = struct{}{}
		}
	case "enter":
		if len(m.selected) == 0 {
			return m, nil
		}

		out := make([]domain.Project, 0, len(m.selected))
		for i := range m.selected {
			out = append(out, m.Projects[i])
		}

		return m, func() tea.Msg { return DoneMsg{Selected: out} }
	}

	return m, nil
}
