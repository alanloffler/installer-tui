package done

import (
	"github.com/alanloffler/installer-tui/internal/installer"
	"github.com/alanloffler/installer-tui/internal/tui/menu"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

type HomeMsg struct{}

type BackMsg struct {
	Next tea.Model
}

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
	Help  key.Binding
	Back  key.Binding
	Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Quit, k.Help}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Enter},
		{k.Back, k.Quit},
		{k.Help},
	}
}

var keys = keyMap{
	Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("up/k", "subir")),
	Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("down/j", "bajar")),
	Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "continuar")),
	Help:  key.NewBinding(key.WithKeys("?", "/"), key.WithHelp("?", "ayuda")),
	Back:  key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "volver")),
	Quit:  key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "salir")),
}

type Model struct {
	Results []installer.Result
	items   []menu.Item
	cursor  int
	help    help.Model
	back    tea.Model
}

func New(results []installer.Result, back tea.Model) Model {
	h := help.New()
	return Model{
		Results: results,
		back:    back,
		help:    h,
		items: []menu.Item{
			{Label: "Volver", Msg: BackMsg{Next: back}},
			{Label: "Inicio", Msg: HomeMsg{}},
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
