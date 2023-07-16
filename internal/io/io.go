package io

import (
	"io"
	"os"

	"github.com/charmbracelet/log"
)

func ReadFile(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Error opening file", "file", file, "error", err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		log.Fatal("Error reading file", "file", file, "error", err)
	}
	return b
}
