package done

import (
	"github.com/alanloffler/bubbletea/internal/installer"

	tea "charm.land/bubbletea/v2"
)

type HomeMsg struct{}
type BackMsg struct{}

type Model struct {
	Results []installer.Result
	cursor  int
}

func New(results []installer.Result) Model {
	return Model{Results: results}
}

func (m Model) Init() tea.Cmd {
	return nil
}
