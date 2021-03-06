# Library for Forge

[![Build Status](https://cloud.drone.io/api/badges/kleister/go-forge/status.svg)](https://cloud.drone.io/kleister/go-forge)
[![Join the Matrix chat at https://matrix.to/#/#kleister:matrix.org](https://img.shields.io/badge/matrix-%23kleister%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#kleister:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/e96f91f1bce14e049a3d3db93baa4683)](https://www.codacy.com/app/kleister/go-forge?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=kleister/go-forge&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/kleister/go-forge?status.svg)](http://godoc.org/github.com/kleister/go-forge)
[![Go Report](http://goreportcard.com/badge/github.com/kleister/go-forge)](http://goreportcard.com/report/github.com/kleister/go-forge)

This repository provides helpers related to Minecraft Forge.


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.11 because we are using Go modules. It is also possible to just simply execute the `go get github.com/kleister/go-forge/...` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/kleister/go-forge/...
cd $GOPATH/src/github.com/kleister/go-forge
make sync gorunpkg clean generate test
```


## Examples

### Fetch available Forge versions

[embedmd]:# (examples/versions/main.go go)
```go
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

	for _, version := range forge.Releases {
		fmt.Println(version.ID)
	}
}
```

### Sort Forge versions by ID

[embedmd]:# (examples/sorted/main.go go)
```go
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
```

### Filter versions by constraints

[embedmd]:# (examples/filter/main.go go)
```go
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
		Version: ">=13.20,<14.23.4",
	}

	for _, version := range forge.Releases.Filter(f) {
		fmt.Println(version.ID)
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
