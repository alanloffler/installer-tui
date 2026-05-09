package installing

import (
	"github.com/alanloffler/bubbletea/internal/domain"
	"github.com/alanloffler/bubbletea/internal/installer"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
)

type DoneMsg struct {
	Results []installer.Result
}

type progressMsg struct {
	Result installer.Result
}

type Model struct {
	Queue    []domain.Project
	Done     []installer.Result
	current  int
	progress progress.Model
	finished bool
}

func New(projects []domain.Project) Model {
	return Model{
		Queue:    projects,
		progress: progress.New(progress.WithDefaultBlend(), progress.WithWidth(40)),
	}
}

func (m Model) Init() tea.Cmd {
	return installAt(m.Queue, 0)
}

func installAt(queue []domain.Project, idx int) tea.Cmd {
	if idx >= len(queue) {
		return nil
	}

	return func() tea.Msg {
		return progressMsg{Result: installer.Install(queue[idx])}
	}
}
