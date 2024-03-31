package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	Motd string `yaml:"motd"`
}

var cfg config

func loadConfig(p string) (err error) {
	var exists bool
	if _, err := os.Stat(p); err == nil {
		exists = true
	} else if errors.Is(err, os.ErrNotExist) {
		exists = false
	} else {
		return err
	}

	cfg = config{}

	if exists {
		data, err := os.ReadFile(p)
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
	err = os.WriteFile(p, data, 0640)
	if err != nil {
		return err
	}

	return nil
}
