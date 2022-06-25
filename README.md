# 🔍 gofind

A command line [pkg.go.dev](pkg.go.dev) searcher

## Usage

Simply specify the keywords in the arguments e.g. package name (`template`), symbol name (`io.Reader`), multiple keywords (`json OR yaml`) etc.

```
Usage:
  gofind [OPTIONS] QUERY...

Application Options:
  -s, --symbol   Search for symbol instead of package
  -l, --limit=   Number of search result items limit (default: 20)
  -V, --version  Show version
  -j, --json     Output in JSON format

Help Options:
  -h, --help     Show this help message
```

To see examples of keywords to search for, check [search-help](https://pkg.go.dev/search-help) on pkg.go.dev.

## Installation

Clone or download this repository then run `go install`, or run below:

```sh
go install github.com/sheepla/gofind@latest
```

## License

MIT

## Author

[Sheepla](https://github.com/sheepla)
