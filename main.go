package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/sheepla/gofind/client"
	"github.com/sheepla/gofind/ui"
)

// nolint:gochecknoglobals
var (
	appName     = "gofind"
	appVersion  = "???"
	appRevision = "???"
	appUsage    = "[OPTIONS] QUERY..."
)

type exitCode int

const (
	exitCodeOK exitCode = iota
	exitCodeErrArgs
	exitCodeErrSearch
	exitCodeErrJSON
	exitCodeFuzzyFinder
	exitCodeErrWebBrowser
)

// nolint:maligned
type options struct {
	SearchForSymbol bool `short:"s" long:"symbol" description:"Search for symbol instead of package"`
	Limit           int  `short:"l" long:"limit" description:"Number of search result items limit" default:"20"`
	Version         bool `short:"V" long:"version" description:"Show version"`
	JSON            bool `short:"j" long:"json" description:"Output in JSON format"`
}

var errNoArgs = errors.New("must require arguments")

func main() {
	code, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(int(code))
}

func run(cliArgs []string) (exitCode, error) {
	var opts options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = appName
	parser.Usage = appUsage

	args, err := parser.ParseArgs(cliArgs)
	if err != nil {
		if flags.WroteHelp(err) {
			return exitCodeOK, nil
		}

		return exitCodeErrArgs, fmt.Errorf("parse error: %w", err)
	}

	if opts.Version {
		// nolint:forbidigo
		fmt.Printf("%s: v%s-rev%s\n", appName, appVersion, appRevision)

		return exitCodeOK, nil
	}

	if len(args) == 0 {
		return exitCodeErrArgs, errNoArgs
	}

	results, err := client.Search(&client.Param{
		Query:           strings.Join(args, " "),
		Limit:           opts.Limit,
		SearchForSymbol: opts.SearchForSymbol,
	})
	if err != nil {
		return exitCodeErrSearch, fmt.Errorf("failed to get search result: %w", err)
	}

	if opts.JSON {
		buf := bufio.NewWriter(os.Stdout)
		if err := json.NewEncoder(buf).Encode(results); err != nil {
			return exitCodeErrJSON, fmt.Errorf("failed to marshalling JSON: %w", err)
		}

		return exitCodeOK, nil
	}

	_, err = ui.Find(*results)
	if err != nil {
		return exitCodeFuzzyFinder, fmt.Errorf("an error occurred on fuzzyfinder: %w", err)
	}

	// if err := ui.Open(&results[idx].Link); err != nil {
	// 	return exitCodeErrWebBrowser, fmt.Errorf("failed to open the link: %w", err)
	// }

	return exitCodeOK, nil
}
