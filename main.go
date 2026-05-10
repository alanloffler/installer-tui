package main

import (
	_ "embed"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/catalog"
	"github.com/alanloffler/bubbletea/internal/tui"
)

//go:embed configs/projects.json
var projectsData []byte

//go:embed configs/packages.json
var packagesData []byte

func main() {
	projects, err := catalog.LoadProjects(projectsData)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	packages, err := catalog.LoadPackages(packagesData)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(tui.NewApp(projects, packages)).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
