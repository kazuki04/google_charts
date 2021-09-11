package config

import (
	"gopkg.in/ini.v1"
	"os"
)

type ConfigList struct {
	Port int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		os.Exit(1)
	}
	Config = ConfigList{Port: cfg.Section("web").Key("port").MustInt()}
}
