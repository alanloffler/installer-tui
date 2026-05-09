package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/alanloffler/bubbletea/internal/catalog"
	"github.com/alanloffler/bubbletea/internal/tui"
)

func main() {
	projects, err := catalog.Load("configs/projects.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(tui.NewApp(projects)).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
