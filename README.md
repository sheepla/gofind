# 🔍 gofind

A command line [pkg.go.dev](pkg.go.dev) searcher

## Usage

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

## Installation

Clone or download this repository then run `go install`, or run below:

```go
go install github.com/sheepla/gofind@latest
```

## License

MIT

## Author

[Sheepla](https://github.com/sheepla)
