package done

import (
	"github.com/alanloffler/bubbletea/internal/installer"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	Results []installer.Result
}

func New(results []installer.Result) Model {
	return Model{Results: results}
}

func (m Model) Init() tea.Cmd {
	return nil
}
