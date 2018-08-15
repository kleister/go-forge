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

	version.ByVersion(forge.Releases).Sort()

	for _, version := range forge.Releases {
		fmt.Println(version.ID)
	}
}
