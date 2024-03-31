package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Motd  string `yaml:"motd"`
	User  string `yaml:"user"`
	Group string `yaml:"group"`
}

var cfg config

var cfgPath string

func loadConfig() (err error) {
	var exists bool
	if _, err := os.Stat(cfgPath); err == nil {
		exists = true
	} else if errors.Is(err, os.ErrNotExist) {
		exists = false
	} else {
		return err
	}

	cfg = config{}

	if exists {
		data, err := os.ReadFile(cfgPath)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(data, cfg)
		if err != nil {
			return err
		}
	}

	if cfg.Motd == "" {
		cfg.Motd = "Hello, World!"
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(cfgPath, data, 0640)
	if err != nil {
		return err
	}

	return nil
}
