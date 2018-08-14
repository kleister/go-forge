# Library for Forge

[![Build Status](http://drone.kleister.tech/api/badges/kleister/go-forge/status.svg)](http://drone.kleister.tech/kleister/go-forge)
[![Stories in Ready](https://badge.waffle.io/kleister/kleister-api.svg?label=ready&title=Ready)](http://waffle.io/kleister/kleister-api)
[![Join the Matrix chat at https://matrix.to/#/#kleister:matrix.org](https://img.shields.io/badge/matrix-%23kleister%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#kleister:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e96f91f1bce14e049a3d3db93baa4683)](https://www.codacy.com/app/kleister/go-forge?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=kleister/go-forge&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/kleister/go-forge?status.svg)](http://godoc.org/github.com/kleister/go-forge)
[![Go Report](http://goreportcard.com/badge/github.com/kleister/go-forge)](http://goreportcard.com/report/github.com/kleister/go-forge)

This repository will provides helpers related to Forge.


## Development

TBD


## Examples

### Fetch available Forge versions

[embedmd]:# (examples/versions/main.go go)
```go
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
```


## Security

If you find a security issue please contact kleister@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
