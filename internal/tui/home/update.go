package home

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, keys.Down):
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case key.Matches(msg, keys.Enter):
			return m, func() tea.Msg { return m.items[m.cursor].Msg }
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.help.SetWidth(msg.Width)
	}

	return m, nil
}
