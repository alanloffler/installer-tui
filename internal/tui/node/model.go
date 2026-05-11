package node

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/installer-tui/internal/domain"
	"github.com/alanloffler/installer-tui/internal/tui/menu"
)

type HomeMsg struct{}

type DoneMsg struct {
	Selected []domain.Package
}

type keyMap struct {
	Up    key.Binding
	Down  key.Binding
	Space key.Binding
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
		{k.Space, k.Enter},
		{k.Back, k.Quit},
		{k.Help},
	}
}

var keys = keyMap{
	Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("up/k", "subir")),
	Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("down/j", "bajar")),
	Space: key.NewBinding(key.WithKeys("space"), key.WithHelp("space", "seleccionar")),
	Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "continuar")),
	Help:  key.NewBinding(key.WithKeys("?", "/"), key.WithHelp("?", "ayuda")),
	Back:  key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "volver")),
	Quit:  key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "salir")),
}

type Model struct {
	Packages []domain.Package
	items    []menu.Item
	cursor   int
	help     help.Model
	selected map[int]struct{}
}

func New(packages []domain.Package) Model {
	h := help.New()
	return Model{
		Packages: packages,
		items:    []menu.Item{{Label: "Volver", Msg: HomeMsg{}}},
		help:     h,
		selected: make(map[int]struct{})}
}

func (m Model) Init() tea.Cmd {
	return nil
}
