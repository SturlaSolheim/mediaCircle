package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Logging  LoggingConfig  `yaml:"logging"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type DatabaseConfig struct {
	Type       string `yaml:"type"`
	Connection string `yaml:"connection"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

var AppConfig *Config

func LoadConfig() error {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	configFile := fmt.Sprintf("config/%s.yaml", env)
	
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("fant ikke config fil: %s", configFile)
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("error ved lesing av config fil: %w", err)
	}

	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return fmt.Errorf("error ved parsing av config fil: %w", err)
	}

	AppConfig = config
	return nil
}

func GetConfigPath(env string) string {
	return filepath.Join("config", fmt.Sprintf("%s.yaml", env))
}

func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}
