<div align="right">

[![Release](https://github.com/sheepla/gofind/actions/workflows/release.yml/badge.svg)](https://github.com/sheepla/gofind/actions/workflows/release.yml)
[![golangci-lint](https://github.com/sheepla/gofind/actions/workflows/ci.yml/badge.svg)](https://github.com/sheepla/gofind/actions/workflows/ci.yml)

</div>

<div align="center">

# 🔍 gofind

*A command line [pkg.go.dev](https://pkg.go.dev) searcher and `go get` helper*

[![MIT](https://img.shields.io/static/v1?label=License&message=MIT&color=blue&style=flat-square)](https://github.com/sheepla/gofind/blob/master/LICENSE)
[![Latest Release](https://img.shields.io/github/v/release/sheepla/gofind?style=flat-square)](https://github.com/sheepla/gofind/releases/latest)

</div>

## Features

- [x] Quickly search Go packages or symbol
- [x] Easily get package that you selected
- [x] Search results can be output in JSON format
- [x] Can be open the document page in your web browser

## Usage

```
Usage:
  gofind [OPTIONS] QUERY...

Application Options:
  -s, --symbol   Search for symbol instead of package
  -l, --limit=   Number of search result items limit (default: 20)
  -V, --version  Show version
  -j, --json     Output search results in JSON format
  -o, --open     Open the document URL in your web browser
  -u, --url      Output pkg.go.dev URL instead of output package name
  -g, --goget    Run go get command to get the package that you selected

Help Options:
  -h, --help     Show this help message
```

Simply specifying the keywords in the arguments e.g. package name (`template`), symbol name (`io.Reader`), multiple keywords (`json OR yaml`) etc.

> **NOTE**:
> To see examples of keywords to search for, check [search-help](https://pkg.go.dev/search-help) on pkg.go.dev.

It can be output the result in JSON format by specifying the `-j`, `--json` option.

By default, it searches for packages, but you can also search for symbols with the `-s`, `--symbol` option.

When you select an item, the name of the package is output. You can also output the URL by specifying the `-u`, `--url` option.

## Installation

You can download the executable binaries from the latest page.

> [![Latest Release](https://img.shields.io/github/v/release/sheepla/gofind?style=flat-square)](https://github.com/sheepla/gofind/releases/latest)

To build from source, clone or download this repository then run `go install`, or run below:

```sh
go install github.com/sheepla/gofind@latest
```

Developing on `go1.18.3 linux/amd64`.

## License

[MIT](./LICENSE)

## Author

[Sheepla](https://github.com/sheepla)
