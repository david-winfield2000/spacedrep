package storage

import (
	"encoding/json"
	"os"
	"spacedrep/models"
)

func LoadConfig() (*models.Config, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &models.Config{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var config models.Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *models.Config) error {
	configFile, err := getConfigFile()
	if err != nil {
		return err
	}

	file, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(config)
}
