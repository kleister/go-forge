# go-forge

[![General Workflow](https://github.com/kleister/go-forge/actions/workflows/general.yml/badge.svg)](https://github.com/kleister/go-forge/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#kleister:matrix.org](https://img.shields.io/badge/matrix-%23kleister%3Amatrix.org-7bc9a4.svg)](https://matrix.to/#/#kleister:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/5ab5e0c9493844ae99f68105629487ed)](https://app.codacy.com/gh/kleister/go-forge/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/kleister/go-forge.svg)](https://pkg.go.dev/github.com/kleister/go-forge)

This repository provides helpers related to Minecraft Forge.

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project doesn't enforce
a specific Go version, but we are trying to use the latest stable releases. It
is also possible to simply execute `go get`, but we prefer to use our `Makefile`:

```console
git clone https://github.com/kleister/go-forge.git
cd go-forge

make clean generate test
```

## Examples

For different kind of examples how to use this library please take a look at the
[examples/](./examples) folder.

## Security

If you find a security issue please contact
[kleister@webhippie.de](mailto:kleister@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```

[golang]: http://golang.org/doc/install.html
