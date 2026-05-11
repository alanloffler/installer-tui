package selector

import (
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/installer-tui/internal/domain"
	"github.com/alanloffler/installer-tui/internal/tui/menu"
)

type HomeMsg struct{}

type DoneMsg struct {
	Selected []domain.Project
}

type Model struct {
	Projects []domain.Project
	items    []menu.Item
	cursor   int
	selected map[int]struct{}
}

func New(projects []domain.Project) Model {
	return Model{
		Projects: projects,
		items:    []menu.Item{{Label: "Volver", Msg: HomeMsg{}}},
		selected: make(map[int]struct{})}
}

func (m Model) Init() tea.Cmd {
	return nil
}
