package main

import (
	"log"

	"github.com/kleister/go-forge/version"
)

func main() {
	log.Println("Fetching Forge versions...")
	forge, err := version.FromDefault()

	if err != nil {
		log.Fatalln(err)
	}

	for _, version := range forge.Releases {
		log.Println("Forge v", version.ID)
	}
}
