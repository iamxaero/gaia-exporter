package config

import (
	"os"

	"github.com/cloudflare/cfssl/log"
	"gopkg.in/yaml.v3"
)

const fileName = "/opt/exporters/gaia-exporter.1.0.0/gaia-exporter/config.yaml"

type Config struct {
	GaiaPort    string `yaml:"Gaia Port"`
	GaiaHost    string `yaml:"Gaia Host"`
	GaiaNetInfo string `yaml:"Gaia Net Info"`
	Interval    int    `yaml:"Interval"`
}

func New() *Config {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Unable to open config file: %v ", err)
	}

	var c Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Debug("Config loaded")
	return &c
}
