package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() *Config {
	data, err := os.ReadFile(getConfigPath())
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return &Config{}
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing json file: ", err)
		return &Config{}
	}

	return &config
}

func (c *Config) SetUser(username string) {
	c.CurrentUserName = username

	err := write(c)
	if err != nil {
		fmt.Println("Error writing file: ", err)
	}
}

func getConfigPath() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
	}

	return usr.HomeDir + "/" + configFileName
}

func write(cfg *Config) error {
	// write the config
	jsonData, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling json: ", err)
		return err
	}

	err = os.WriteFile(getConfigPath(), jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
		return err
	}

	return nil
}
