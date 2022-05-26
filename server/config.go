package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type conf struct {
	jrzd_dir string
	xd_dir   string
	zip_dir  string
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
	fmt.Println(path)
	if _, err := toml.DecodeFile(path, &cf); err != nil {
		print(err)
		return err
	}

	return nil
}
