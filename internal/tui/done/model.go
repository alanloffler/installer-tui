package done

import (
	"github.com/alanloffler/bubbletea/internal/installer"
	"github.com/alanloffler/bubbletea/internal/tui/menu"

	tea "charm.land/bubbletea/v2"
)

type HomeMsg struct{}

type BackMsg struct {
	Next tea.Model
}

type Model struct {
	Results []installer.Result
	items   []menu.Item
	cursor  int
	back    tea.Model
}

func New(results []installer.Result, back tea.Model) Model {
	return Model{
		Results: results,
		back:    back,
		items: []menu.Item{
			{Label: "Volver", Msg: BackMsg{Next: back}},
			{Label: "Inicio", Msg: HomeMsg{}},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
