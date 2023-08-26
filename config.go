package main

import (
	"errors"
	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Motd string `yaml:"motd"`
}

var cfg *Config

func init() {
	loaded, err := loadConfig()
	if err != nil {
		log.Fatalf("error loading config: %s\n", err.Error())
	}
	cfg = loaded
}

func loadConfig() (cfg *Config, err error) {
	path, err := xdg.ConfigFile("janitord.yaml")
	if err != nil {
		return nil, err
	}

	var exists bool
	if _, err := os.Stat(path); err == nil {
		exists = true
	} else if errors.Is(err, os.ErrNotExist) {
		exists = false
	} else {
		return nil, err
	}

	switch exists {
	case true:
		cfg = &Config{}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, cfg)
		if err != nil {
			return nil, err
		}
	case false:
		cfg = &Config{Motd: "Hello, World!"}
		data, err := yaml.Marshal(cfg)
		if err != nil {
			return nil, err
		}
		err = os.WriteFile(path, data, 0640)
		if err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
