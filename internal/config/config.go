package config

import (
	"os"
	"log"
	"fmt"
	"encoding/json"
)

type Config struct {
	URL					string			`json:"db_url"`
	UserName		string			`json:"current_user_name"`
}

func ReadConfigJson() (Config, error) {
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

func (c *Config) SetUser(currentUserName string) error {
		c.UserName = currentUserName

		jsonData, err := json.MarshalIndent(c, "", "	")
    if err != nil {
      log.Fatal(err)
    }

		filePath, err := getConfigFilePath()
		if err != nil {
			return err
		}

		if err = os.WriteFile(filePath, jsonData, 0644); err != nil {
			return fmt.Errorf("write config: %w", err)
		}

		return nil

}
