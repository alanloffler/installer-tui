package home

import tea "charm.land/bubbletea/v2"

type menuItem struct {
	label string
	msg   tea.Msg
}

type GoToSelectorMsg struct{}
type GoToNodePkgMsg struct{}

type Model struct {
	items  []menuItem
	cursor int
}

func New() Model {
	return Model{
		items: []menuItem{
			{label: "Instalar proyectos", msg: GoToSelectorMsg{}},
			{label: "Instalar paquetes node", msg: GoToNodePkgMsg{}},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
