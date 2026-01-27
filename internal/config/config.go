package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	if err != nil {
		log.Println("Error getting config file path:", err)
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		log.Println("Error opening JSON config:", err)
		return Config{}, nil
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		log.Println("Error decoding json config file:", err)
		return Config{}, nil
	}

	return config, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name

	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, configFileName), nil
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		log.Println("Error getting config file path", err)
		return err
	}
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
