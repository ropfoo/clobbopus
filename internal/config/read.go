package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Port    int `yaml:"port"`
	Domains map[string]struct {
		Queries []string `yaml:"queries"`
		Url     string   `yaml:"url"`
	} `yaml:"domains"`
}

func Read() Config {
	yfile, err := os.ReadFile("../clobbopus.yml")
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err2 := yaml.Unmarshal(yfile, &config)
	if err2 != nil {
		log.Fatal(err2)
	}

	return config
}
