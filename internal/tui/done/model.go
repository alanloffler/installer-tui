package done

import (
	"github.com/alanloffler/bubbletea/internal/installer"

	tea "charm.land/bubbletea/v2"
)

type HomeMsg struct{}

type BackMsg struct {
	Next tea.Model
}

type Model struct {
	Results []installer.Result
	cursor  int
	back    tea.Model
}

func New(results []installer.Result, back tea.Model) Model {
	return Model{Results: results, back: back}
}

func (m Model) Init() tea.Cmd {
	return nil
}
