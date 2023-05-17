package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Domains map[string]struct {
		Params map[string]string `yaml:"params"`
		Urls   []string          `yaml:"urls"`
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

	fmt.Println(config.Domains["musikersucht"].Urls)

	return config
}
