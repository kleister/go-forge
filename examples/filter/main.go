package main

import (
	"fmt"
	"os"

	"github.com/kleister/go-forge/version"
)

func main() {
	fmt.Println("Fetching Forge versions...")
	forge, err := version.FromDefault()

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	f := &version.Filter{
		Version:   ">=13.20,<14.23.4",
		Minecraft: ">1.12",
	}

	for _, version := range forge.Releases.Filter(f) {
		fmt.Println(version.ID)
	}
}
