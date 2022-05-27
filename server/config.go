package main

import (
	"github.com/BurntSushi/toml"
)

type conf struct {
	Jrzd_dir string
	Xd_dir   string
	Zip_dir  string
}

var cf conf

func load_default_config() error {
	var path string = "./conf.toml"
	if _, err := toml.DecodeFile(path, &cf); err != nil {
		return err
	}

	return nil
}

func load_config(path string) error {
	if _, err := toml.DecodeFile(path, &cf); err != nil {
		return err
	}
	return nil
}
