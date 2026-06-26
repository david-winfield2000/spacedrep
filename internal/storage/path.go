package storage

import (
	"os"
	"path/filepath"
)

func appDataPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(dir, "spacedrep")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return "", err
	}

	return appDir, nil
}

func gitRepoPath() (string, error) {
	dir, err := appDataPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "repo"), nil
}

func getConfigFile() (string, error) {
	dir, err := appDataPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "config.json"), nil
}

func getProblemsFile() (string, error) {
	dir, err := gitRepoPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "problems.json"), nil
}
