package config

import (
	"os"

	"github.com/cloudflare/cfssl/log"
	"gopkg.in/yaml.v3"
)

const CI = "gitlab"
const fileName = "config.yaml"

type Config struct {
	GaiaBin string   `yaml:"Gaia Bin Path"`
	GaiaPort string     `yaml:"Gaia Port"`
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
