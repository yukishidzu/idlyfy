package main

import (
	"log"

	"github.com/yukishidzu/idlyfy/config"
	"github.com/yukishidzu/idlyfy/ui"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	ui.Run()
}
