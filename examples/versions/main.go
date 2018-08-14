package main

import (
	"log"

	"github.com/kleister/go-forge/version"
)

func main() {
	log.Println("Fetching Forge versions...")
	forge, err := version.FromDefault()

	if err != nil {
		log.Fatalf("%s", err)
	}

	for _, version := range forge.Releases {
		log.Printf("Forge v%s", version.ID)
	}
}
