package done

import (
	tea "charm.land/bubbletea/v2"
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
		if m.cursor < 1 {
			m.cursor++
		}
	case "enter":
		switch m.cursor {
		case 0:
			return m, func() tea.Msg { return BackMsg{Next: m.back} }
		case 1:
			return m, func() tea.Msg { return HomeMsg{} }
		}
	}

	return m, nil
}
