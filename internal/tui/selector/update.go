package selector

import (
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/installer-tui/internal/domain"
	"github.com/alanloffler/installer-tui/internal/tui/browser"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return m, nil
	}

	projSlots := cursorSlotsPerProj * len(m.Projects)
	maxCursor := projSlots + len(m.items) - 1

	switch key.String() {
	case "ctrl-c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < maxCursor {
			m.cursor++
		}
	case "space":
		if m.cursor >= projSlots || m.cursor%cursorSlotsPerProj == 1 {
			return m, nil
		}

		idx := m.cursor / cursorSlotsPerProj
		if _, ok := m.selected[idx]; ok {
			delete(m.selected, idx)
		} else {
			m.selected[idx] = struct{}{}
		}
	case "enter":
		if m.cursor >= projSlots {
			idx := m.cursor - projSlots
			return m, func() tea.Msg { return m.items[idx].Msg }
		}

		idx := m.cursor / cursorSlotsPerProj
		if m.cursor%cursorSlotsPerProj == 1 {
			p := m.Projects[idx]
			if p.Repo != "" {
				browser.Open(p.Repo)
			}

			return m, nil
		}

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
