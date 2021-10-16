package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Application configuration.
type AppConfig struct {
	Port    string         `yaml:"port"`
	Devices []DeviceConfig `yaml:"devices"`
}

// Device configuration.
type DeviceConfig struct {
	Type       string      `yaml:"type"`
	MacAddress string      `yaml:"mac_address"`
	Tags       []TagConfig `yaml:"tags"`
}

// Tag configuration
type TagConfig struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Reads the application configuration from a file.
func readAppConfig(filename string) (*AppConfig, error) {
	config := &AppConfig{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
