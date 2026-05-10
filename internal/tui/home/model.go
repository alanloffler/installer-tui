package home

import (
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/tui/menu"
)

type GoToSelectorMsg struct{}
type GoToNodePkgMsg struct{}

type Model struct {
	items  []menu.Item
	cursor int
}

func New() Model {
	return Model{
		items: []menu.Item{
			{Label: "Template de proyectos", Msg: GoToSelectorMsg{}},
			{Label: "Paquetes de node", Msg: GoToNodePkgMsg{}},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
