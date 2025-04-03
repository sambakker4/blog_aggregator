package config

import (
	"encoding/json"
	"os"
)

const configFilename = "/.gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

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
