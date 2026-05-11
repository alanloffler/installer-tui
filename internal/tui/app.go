package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/domain"
	"github.com/alanloffler/bubbletea/internal/tui/done"
	"github.com/alanloffler/bubbletea/internal/tui/header"
	"github.com/alanloffler/bubbletea/internal/tui/home"
	"github.com/alanloffler/bubbletea/internal/tui/installing"
	"github.com/alanloffler/bubbletea/internal/tui/node"
	"github.com/alanloffler/bubbletea/internal/tui/selector"
)

const appTitle = "📦 Instalador"
const appSubtitle = "0.0.1"

type subtitler interface {
	SectionSubtitle() string
}

type App struct {
	current  tea.Model
	projects []domain.Project
	packages []domain.Package
}

func NewApp(projects []domain.Project, packages []domain.Package) App {
	return App{current: home.New(), projects: projects, packages: packages}
}

func (a App) Init() tea.Cmd {
	return a.current.Init()
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m := msg.(type) {
	case home.GoToSelectorMsg:
		a.current = selector.New(a.projects)
		return a, a.current.Init()
	case home.GoToNodePkgMsg:
		a.current = node.New(a.packages)
		return a, a.current.Init()
	case selector.HomeMsg:
		a.current = home.New()
		return a, a.current.Init()
	case selector.DoneMsg:
		back := selector.New(a.projects)
		a.current = installing.NewFromProjects(m.Selected, back)
		return a, a.current.Init()
	case node.HomeMsg:
		a.current = home.New()
		return a, a.current.Init()
	case node.DoneMsg:
		back := node.New(a.packages)
		a.current = installing.NewFromPackages(m.Selected, back)
		return a, a.current.Init()
	case installing.DoneMsg:
		a.current = done.New(m.Results, m.Next)
		return a, a.current.Init()
	case done.BackMsg:
		a.current = m.Next
		return a, a.current.Init()
	case done.HomeMsg:
		a.current = home.New()
		return a, a.current.Init()
	}

	next, cmd := a.current.Update(msg)
	a.current = next

	return a, cmd
}

func (a App) View() tea.View {
	inner := a.current.View()
	sectionSubtitle := ""

	if s, ok := a.current.(subtitler); ok {
		sectionSubtitle = s.SectionSubtitle()
	}

	inner.Content = header.Render(appTitle, appSubtitle, sectionSubtitle) + inner.Content
	inner.AltScreen = true

	return inner
}
