package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
		SSLMode  string `json:"sslmode"`
	} `json:"database"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func ReadServiceConfig(service string) (string, error) {
	configFile := "config/config.json"
	absConfigPath, err := filepath.Abs(configFile)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	}
	viper.SetConfigFile(absConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}

	serviceConfig := "service_url." + service
	serviceURL := viper.GetString(serviceConfig)

	return serviceURL, nil
}

func ReadAPIConfig(api string) (string, error) {
	configFile := "config/config.json"
	absConfigPath, err := filepath.Abs(configFile)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	}
	viper.SetConfigFile(absConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}

	apiConfig := "api_routes." + api
	apiURL := viper.GetString(apiConfig)

	return apiURL, nil
}
