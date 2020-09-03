package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sergalkin/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalf("Cannot parse configuration file at %s. Error=%s", configPath, err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}