package config

import (
	"os"
	"encoding/json"
)

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(path)

	if err != nil {
		return Config{}, err
	}
	var contents Config

	err = json.Unmarshal(data, &contents)
	if err != nil {
		return Config{}, err
	}
	return contents, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homePath + configFilename, nil
}
