package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Configuration file name stored in the home directory
const configFileName = ".gatorconfig.json"

// Config struct holds database URL and current username
type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// SetUser updates the current username and writes it to .gatorconfig.json
func (cfg *Config) SetUser(name string) error {
	cfg.CurrentUserName = name
	// Calls the helper function to write changes to the config file
	return write(*cfg)
}

// Read reads from the config file, decodes it into a Config struct, and returns it
func Read() (Config, error) {
	// Get the full path of the config file
	configPath, err := GetConfigPath()
	if err != nil {
		return Config{}, err
	}
	// Open the file at the retrieved path, returns a file descriptor
	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	// Close the file when the function exits
	defer file.Close()

	// Decode JSON data from the file into the Config struct
	cfg := Config{}
	decode := json.NewDecoder(file)
	err = decode.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

// GetConfigPath returns the full path of the configuration file
func GetConfigPath() (string, error) {
	// Get the home directory path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Join the home directory with the config file name to get the full path
	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}

// Helper function that writes the Config struct data to the config file
func write(cfg Config) error {
	// Get the config file path
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	// Create or overwrite the config file, returns a file descriptor
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	// Ensure the file is closed when the function exits
	defer file.Close()

	// Encode and write struct data as JSON to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}
