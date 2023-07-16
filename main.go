package main

import (
	"github.com/charmbracelet/log"
	"github.com/ivanvc/tpl/internal/config"
	"github.com/ivanvc/tpl/internal/template"
)

func main() {
	cfg := config.Load()
	if cfg.SetDebugLogLevel {
		log.Default().SetLevel(log.DebugLevel)
	}
	log.Debug("Loaded configuration", "config", cfg)

	template.Run(cfg)
}
