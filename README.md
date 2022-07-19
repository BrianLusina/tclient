# TClient

[![Build](https://github.com/BrianLusina/tclient/actions/workflows/build_app.yml/badge.svg)](https://github.com/BrianLusina/tclient/actions/workflows/build_app.yml)
[![Lint](https://github.com/BrianLusina/tclient/actions/workflows/lint.yml/badge.svg)](https://github.com/BrianLusina/tclient/actions/workflows/lint.yml)
[![Tests](https://github.com/BrianLusina/tclient/actions/workflows/tests.yml/badge.svg)](https://github.com/BrianLusina/tclient/actions/workflows/tests.yml)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/0aa1d4aa20d84894bb6b23524e0a3956)](https://www.codacy.com/gh/BrianLusina/tclient/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=BrianLusina/tclient&amp;utm_campaign=Badge_Grade)

A simple BitTorrent client in Go.

## Getting started

This client is pretty small & simple, however it has been structured to use minimal dependencies and follows a standard pattern to make the code easier to read.

There are a couple of things needed to get started.

1. [Go 1.18](https://go.dev/)

Instructions on installing Go 1.18 can be found in the link provided.

Pretty much that's it :)

Now, download dependencies with:

```bash
make install
```

## Usage

The options available are:

``` plain
Usage:
  -input string
        Input file
  -output string
        Where to download the file (default "./")
```

You can build the application and run it or run it directly with:

``` bash
go run cmd/cli/main.go -input <TORRENT_FILE_PATH> -output <TORRENT_FILE_OUTPUT>
```

`<TORRENT_FILE_PATH>` is the file path of the torrent file, `<TORRENT_FILE_OUTPUT>` is the output path that you intend the file to be downloaded to.

For example:

```bash
go run cmd/cli/main.go -input debian-11.4.0-amd64-netinst.iso.torrent -output debian.iso
```

Optionally, you can build the application

```bash
make build
# or
go build -o <OUTPUT_PATH> cmd/cli/main.go
```

> This will build the application in the [bin](./bin/) directory. If you use the second option to build the application, you can specify the `<OUTPUT_PATH>` directory.

Once done, you can use the binary as well:

```bash
tclient -input debian-10.2.0-amd64-netinst.iso.torrent -output debian.iso
```

> This assumes the binary built is called `tclient`

## Tests

Tests have been written for the source code and can be validated with running:

```bash
make test
# or
go test ./...
```

## Linting

Linting can be done using [golanci-lint](https://golangci-lint.run/) and setup is as below:

```bash
make setup-linting
```

Now run linting with:

```bash
make lint
```

Other commands can be found in the [Makefile](./Makefile).

## Limitations

- Only supports .torrent files (no magnet links)
- Only supports HTTP trackers
- Does not support multi-file torrents
- Strictly leeches (does not support uploading pieces)

## Versioning

[SemVer](https://semver.org/) is used for versioning. For the versions available, see the [tags](https://github.com/BrianLusina/tclient/tags) in this repository.

## License

View the project license [here](./LICENSE)
