package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config holds the application configuration
type Config struct {
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}
}

// LoadConfig loads the configuration from a file
func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
