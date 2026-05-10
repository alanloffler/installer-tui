package home

import tea "charm.land/bubbletea/v2"

type GoToSelectorMsg struct{}
type GoToNodePkgMsg struct{}

type menuItem struct {
	label string
	msg   tea.Msg
}

type Model struct {
	items  []menuItem
	cursor int
}

func New() Model {
	return Model{
		items: []menuItem{
			{label: "Template de proyectos", msg: GoToSelectorMsg{}},
			{label: "Paquetes de node", msg: GoToNodePkgMsg{}},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
