package installer

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/alanloffler/installer-tui/internal/domain"
)

type Result struct {
	Name      string
	Version   string
	Output    string
	PkgMan    string
	Err       error
	UsageHint string
}

type packageJSON struct {
	PackageManager string `json:"packageManager"`
	DevEngines     *struct {
		PackageManager *struct {
			Name string `json:"name"`
		} `json:"packageManager"`
	} `json:"devEngines"`
}

type locks struct {
	file string
	pm   string
}

func DetectPM() string {
	locks := []locks{
		{"pnpm-lock.yaml", "pnpm"},
		{"yarn.lock", "yarn"},
		{"bun.lock", "bun"},
		{"package-lock.json", "npm"},
	}

	for _, l := range locks {
		if _, err := os.Stat(l.file); err == nil {
			return l.pm
		}
	}

	data, err := os.ReadFile("package.json")
	if err != nil {
		return ""
	}

	var pkg packageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return ""
	}

	if pkg.PackageManager != "" {
		name, _, _ := strings.Cut(pkg.PackageManager, "@")
		return name
	}

	if pkg.DevEngines != nil && pkg.DevEngines.PackageManager != nil {
		return pkg.DevEngines.PackageManager.Name
	}

	return ""
}

var ErrAlreadyExists = errors.New("instalado previamente")
var ErrPackageNotFound = errors.New("paquete no encontrado en el registro")
var ErrNoProject = errors.New("no hay proyecto inicializado en este directorio")

func InstallProject(p domain.Project) Result {
	out, err := exec.Command("sh", "-c", p.InstallCmd).CombinedOutput()
	output := string(out)

	r := Result{Name: p.Name, Output: output, UsageHint: p.UsageHint}
	if err != nil {
		r.Err = classify(output, err)
	}

	return r
}

func InstallPackage(p domain.Package) Result {
	if _, err := os.Stat("package.json"); err != nil {
		return Result{Name: p.Name, Err: ErrNoProject}
	}

	pm := DetectPM()
	if pm == "" {
		pm = "npm"
	}

	cmd := p.InstallCmd
	if cmd == "" {
		installCmd := "install"
		if pm != "npm" {
			installCmd = "add"
		}
		cmd = pm + " " + installCmd + " " + p.Cmd
	}

	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	output := string(out)

	r := Result{Name: p.Name, PkgMan: pm, Output: output}
	if err != nil {
		r.Err = classify(output, err)
		return r
	}

	viewCmd := "npm view " + p.Cmd + " version"
	switch pm {
	case "bun":
		viewCmd = "bun info " + p.Cmd + " version"
	case "pnpm":
		viewCmd = "pnpm view " + p.Cmd + " version"
	case "yarn":
		viewCmd = "yarn info " + p.Cmd + " version"
	}

	if v, err := exec.Command("sh", "-c", viewCmd).Output(); err == nil {
		r.Version = strings.TrimSpace(string(v))
	}

	return r
}

func classify(output string, err error) error {
	low := strings.ToLower(output)

	switch {
	case strings.Contains(low, "already exists"),
		strings.Contains(low, "already installed"):
		return ErrAlreadyExists
	case strings.Contains(low, "e404"),
		strings.Contains(low, "404"),
		strings.Contains(low, "not found"),
		strings.Contains(low, "not in this registry"),
		strings.Contains(low, "enoent"):
		return ErrPackageNotFound
	}

	if msg := strings.TrimSpace(output); msg != "" {
		return fmt.Errorf("%s (%w)", msg, err)
	}

	return err
}
