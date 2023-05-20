package config

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Port    int `yaml:"port"`
	Domains map[string]struct {
		Params []string `yaml:"params"`
		Url    string   `yaml:"url"`
	} `yaml:"domains"`
}

func (c *Config) Route(domain string) string {
	var route string
	url := c.Domains[domain].Url
	routeSection := strings.SplitN(url, "/", 4)[3]
	route = "/" + domain + "/" + routeSection
	return route
}

func (c *Config) Subroute(domain string) string {
	var route string
	url := c.Domains[domain].Url
	route = strings.SplitN(url, "/", 4)[3]
	return route
}

func Read() Config {
	yfile, err := os.ReadFile("../../clobbopus.yml")
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
