package node

import (
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/domain"
	"github.com/alanloffler/bubbletea/internal/tui/menu"
)

type HomeMsg struct{}

type DoneMsg struct {
	Selected []domain.Package
}

type Model struct {
	Packages []domain.Package
	items    []menu.Item
	cursor   int
	selected map[int]struct{}
}

func New(packages []domain.Package) Model {
	return Model{
		Packages: packages,
		items:    []menu.Item{{Label: "Volver", Msg: HomeMsg{}}},
		selected: make(map[int]struct{})}
}

func (m Model) Init() tea.Cmd {
	return nil
}
