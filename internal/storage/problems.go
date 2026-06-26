package storage

import (
	"encoding/json"
	"os"
	"spacedrep/models"
)

func SaveProblems(problems []*models.Problem) error {
	data, err := json.MarshalIndent(problems, "", "  ")
	if err != nil {
		return err
	}

	path, err := getProblemsFile()
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadProblems() ([]*models.Problem, error) {
	path, err := getProblemsFile()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []*models.Problem{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []*models.Problem{}, nil
	}

	var problems []*models.Problem
	if err := json.Unmarshal(data, &problems); err != nil {
		return nil, err
	}

	return problems, nil
}
