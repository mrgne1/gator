package config

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	User             string `json:"current_user_name"`
	ConnectionString string `json:"connection_string"`
}

var configName string = ".gatorconfig.json"

func Read() (Config, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	configPath := path.Join(homePath, configName)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(name string) error {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := path.Join(homePath, configName)

	c.User = name

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(configPath, data, 0666)
	return err
}
