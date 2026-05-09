package installer

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/alanloffler/bubbletea/internal/domain"
)

var ErrAlreadyExists = errors.New("ya estaba instalado en esta carpeta")

type Result struct {
	Project domain.Project
	Output  string
	Err     error
}

func Install(p domain.Project) Result {
	out, err := exec.Command("sh", "-c", p.InstallCmd).CombinedOutput()

	output := string(out)
	if err == nil {
		return Result{Project: p, Output: output}
	}

	return Result{Project: p, Output: string(out), Err: classify(output, err)}
}

func classify(output string, err error) error {
	low := strings.ToLower(output)

	switch {
	case strings.Contains(low, "already exists"),
		strings.Contains(low, "already installed"):
		return ErrAlreadyExists
	}

	if msg := strings.TrimSpace(output); msg != "" {
		return fmt.Errorf("%s (%w)", msg, err)
	}

	return err
}
