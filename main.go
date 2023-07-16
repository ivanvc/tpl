package main

import (
	"github.com/ivanvc/tpl/internal/config"
	"github.com/ivanvc/tpl/internal/template"
)

func main() {
	cfg := config.Load()
	template.Run(cfg)
}
