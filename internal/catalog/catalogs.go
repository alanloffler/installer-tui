package catalog

import (
	"encoding/json"
	"fmt"

	"github.com/alanloffler/installer-tui/internal/domain"
)

func LoadProjects(data []byte) ([]domain.Project, error) {
	var projects []domain.Project
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, fmt.Errorf("parse projects catalog error: %w", err)
	}

	return projects, nil
}

func LoadPackages(data []byte) ([]domain.Package, error) {
	var pkgs []domain.Package
	if err := json.Unmarshal(data, &pkgs); err != nil {
		return nil, fmt.Errorf("parse packages catalog error: %w", err)
	}

	return pkgs, nil
}
