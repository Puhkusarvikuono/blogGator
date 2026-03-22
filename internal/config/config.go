package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	newConfig := Config{}

	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}

	dat, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	if err := json.Unmarshal(dat, &newConfig); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	return newConfig, nil
}

func getConfigFilePath() (string, error) {
	const configFileName = "/.gatorconfig.json"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + configFileName, nil
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrentUserName = userName

	jsonData, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err = os.WriteFile(filePath, jsonData, 0o644); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}
