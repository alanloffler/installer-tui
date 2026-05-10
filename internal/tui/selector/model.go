package selector

import (
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/domain"
)

type HomeMsg struct{}

type DoneMsg struct {
	Selected []domain.Project
}

type Model struct {
	Projects []domain.Project
	cursor   int
	selected map[int]struct{}
}

func New(projects []domain.Project) Model {
	return Model{Projects: projects, selected: make(map[int]struct{})}
}

func (m Model) Init() tea.Cmd {
	return nil
}
