package installing

import (
	"github.com/alanloffler/bubbletea/internal/domain"
	"github.com/alanloffler/bubbletea/internal/installer"
	"github.com/alanloffler/bubbletea/internal/tui/styles"

	"charm.land/bubbles/v2/progress"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
)

type DoneMsg struct {
	Results []installer.Result
	Next    tea.Model
}

type progressMsg struct {
	Result installer.Result
}

type Job struct {
	Name    string
	Install func() installer.Result
}

type Model struct {
	Queue    []Job
	Done     []installer.Result
	current  int
	progress progress.Model
	spinner  spinner.Model
	finished bool
	PM       string
	back     tea.Model
}

func NewFromProjects(projects []domain.Project, back tea.Model) Model {
	jobs := make([]Job, len(projects))

	for i, p := range projects {
		jobs[i] = Job{Name: p.Name, Install: func() installer.Result { return installer.InstallProject(p) }}
	}

	return newWithJobs(jobs, back)
}

func NewFromPackages(pkgs []domain.Package, back tea.Model) Model {
	jobs := make([]Job, len(pkgs))

	for i, p := range pkgs {
		jobs[i] = Job{Name: p.Name, Install: func() installer.Result { return installer.InstallPackage(p) }}
	}

	return newWithJobs(jobs, back)
}

func newWithJobs(jobs []Job, back tea.Model) Model {
	s := spinner.New(
		spinner.WithSpinner(spinner.MiniDot),
		spinner.WithStyle(styles.WarningStyle),
	)

	return Model{
		Queue: jobs,
		back:  back,
		PM:    installer.DetectPM(),
		progress: progress.New(
			progress.WithColors(styles.ColorYellow, styles.ColorYellow),
			progress.WithWidth(40),
		),
		spinner: s,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(installAt(m.Queue, 0), m.spinner.Tick)
}

func installAt(queue []Job, idx int) tea.Cmd {
	if idx >= len(queue) {
		return nil
	}

	return func() tea.Msg {
		return progressMsg{Result: queue[idx].Install()}
	}
}
