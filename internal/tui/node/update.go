package node

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/installer-tui/internal/domain"
	"github.com/alanloffler/installer-tui/internal/tui/browser"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	pkgSlots := cursorSlotsPerPkg * len(m.Packages)
	maxCursor := pkgSlots + len(m.items) - 1

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {

		case key.Matches(msg, keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, keys.Down):
			if m.cursor < maxCursor {
				m.cursor++
			}
		case key.Matches(msg, keys.Space):
			if m.cursor >= pkgSlots || m.cursor%cursorSlotsPerPkg == 1 {
				return m, nil
			}

			idx := m.cursor / cursorSlotsPerPkg
			if _, ok := m.selected[idx]; ok {
				delete(m.selected, idx)
			} else {
				m.selected[idx] = struct{}{}
			}
		case key.Matches(msg, keys.Enter):
			if m.cursor >= pkgSlots {
				idx := m.cursor - pkgSlots
				return m, func() tea.Msg { return m.items[idx].Msg }
			}

			idx := m.cursor / cursorSlotsPerPkg
			if m.cursor%cursorSlotsPerPkg == 1 {
				p := m.Packages[idx]
				if p.Repo != "" {
					browser.Open(p.Repo)
				}

				return m, nil
			}

			if len(m.selected) == 0 {
				return m, nil
			}

			out := make([]domain.Package, 0, len(m.selected))
			for i := range m.selected {
				out = append(out, m.Packages[i])
			}

			return m, func() tea.Msg { return DoneMsg{Selected: out} }
		case key.Matches(msg, keys.Back):
			return m, func() tea.Msg { return HomeMsg{} }
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	}

	return m, nil
}
