package home

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/installer-tui/internal/tui/menu"
)

type GoToSelectorMsg struct{}
type GoToNodePkgMsg struct{}

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Enter, k.Quit},
		{k.Help},
	}
}

var keys = keyMap{
	Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("up/k", "subir")),     //↑
	Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("down/j", "bajar")), //↓
	Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "seleccionar")),
	Help:  key.NewBinding(key.WithKeys("?", "/"), key.WithHelp("?", "ayuda")),
	Quit:  key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "salir")),
}

type Model struct {
	items  []menu.Item
	cursor int
	help   help.Model
}

func New() Model {
	h := help.New()
	return Model{
		help: h,
		items: []menu.Item{
			{Label: "Template de proyectos", Msg: GoToSelectorMsg{}},
			{Label: "Paquetes de node", Msg: GoToNodePkgMsg{}},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
