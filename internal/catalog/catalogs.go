package catalog

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alanloffler/bubbletea/internal/domain"
)

func Load(path string) ([]domain.Project, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read catalog error: %w", err)
	}

	var projects []domain.Project
	if err := json.Unmarshal(data, &projects); err != nil {
		return nil, fmt.Errorf("parse catalog error: %w", err)
	}

	return projects, nil
}
