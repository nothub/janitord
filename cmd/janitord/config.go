package main

import (
	"errors"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

type config struct {
	Motd string `yaml:"motd"`
}

func loadConfig() (cfg *config, err error) {
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

	cfg = &config{}

	if exists {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, cfg)
		if err != nil {
			return nil, err
		}
	}

	if cfg.Motd == "" {
		cfg.Motd = "Hello, World!"
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(path, data, 0640)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
