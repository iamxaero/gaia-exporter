package controller

import (
	"example.com/gaia-exporter/config"
)

type Controller struct {
	Config *config.Config
}

func New(config *config.Config) *Controller {
	return &Controller{
		Config: config,
	}
}
