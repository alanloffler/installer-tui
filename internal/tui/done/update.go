package done

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok {
		switch key.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "b", "esc":
			return m, func() tea.Msg { return BackMsg{} }
		case "h":
			return m, func() tea.Msg { return HomeMsg{} }
		}
	}

	return m, nil
}
