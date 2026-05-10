package installing

import (
	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.finished {
				results := m.Done
				return m, func() tea.Msg { return DoneMsg{Results: results, Next: m.back} }
			}
		}
	case progressMsg:
		m.Done = append(m.Done, msg.Result)
		m.current++
		pct := float64(m.current) / float64(len(m.Queue))
		cmds := []tea.Cmd{m.progress.SetPercent(pct)}

		if m.current < len(m.Queue) {
			cmds = append(cmds, installAt(m.Queue, m.current))
		}

		return m, tea.Batch(cmds...)
	case progress.FrameMsg:
		pm, cmd := m.progress.Update(msg)
		m.progress = pm

		if !m.finished && m.current >= len(m.Queue) && !m.progress.IsAnimating() {
			m.finished = true
		}

		return m, cmd
	}

	return m, nil
}
