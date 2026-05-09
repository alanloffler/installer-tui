package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/domain"
	"github.com/alanloffler/bubbletea/internal/tui/done"
	"github.com/alanloffler/bubbletea/internal/tui/header"
	"github.com/alanloffler/bubbletea/internal/tui/home"
	"github.com/alanloffler/bubbletea/internal/tui/installing"
	"github.com/alanloffler/bubbletea/internal/tui/selector"
)

const appTitle = "📦 Project Installer"
const appSubtitle = "v0.0.1"

type subtitler interface {
	SectionSubtitle() string
}

type App struct {
	current  tea.Model
	projects []domain.Project
}

func NewApp(projects []domain.Project) App {
	return App{current: home.New(), projects: projects}
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
		return a, nil
	case selector.DoneMsg:
		a.current = installing.New(m.Selected)
		return a, a.current.Init()
	case installing.DoneMsg:
		a.current = done.New(m.Results)
		return a, a.current.Init()
	case done.BackMsg:
		a.current = selector.New(a.projects)
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
