package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/ktr0731/go-fuzzyfinder"
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
	exitCodeErrFuzzyFinder
	exitCodeErrWebBrowser
	exitCodeErrGoGet
)

// nolint:maligned
type options struct {
	SearchForSymbol bool `short:"s" long:"symbol" description:"Search for symbol instead of package"`
	Limit           int  `short:"l" long:"limit" description:"Number of search result items limit" default:"20"`
	Version         bool `short:"V" long:"version" description:"Show version"`
	JSON            bool `short:"j" long:"json" description:"Output search results in JSON format"`
	Open            bool `short:"o" long:"open" description:"Open the document URL in your web browser"`
	URL             bool `short:"u" long:"url" description:"Output pkg.go.dev URL instead of output package name"`
	GoGet           bool `short:"g" long:"goget" description:"Run go get command to get the package that you selected"`
}

var (
	errNoArgs       = errors.New("must require arguments")
	errConflictOpts = errors.New("this option cannot be used at the same time")
)

func main() {
	code, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(int(code))
}

// nolint:cyclop,funlen
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

	if opts.SearchForSymbol && opts.URL {
		return exitCodeErrArgs, fmt.Errorf("%w (--symbol, --url)", errConflictOpts)
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
		if err := json.NewEncoder(os.Stdout).Encode(results); err != nil {
			return exitCodeErrJSON, fmt.Errorf("failed to marshalling JSON: %w", err)
		}

		return exitCodeOK, nil
	}

	idx, err := ui.Find(results)
	if err != nil {
		if errors.Is(fuzzyfinder.ErrAbort, err) {
			return exitCodeOK, nil
		}

		return exitCodeErrFuzzyFinder, fmt.Errorf("an error occurred on fuzzyfinder: %w", err)
	}

	if opts.Open {
		if err := ui.Open(results[idx].Link); err != nil {
			return exitCodeErrWebBrowser, fmt.Errorf("failed to open the link: %w", err)
		}
	}

	if opts.URL {
		fmt.Fprintln(os.Stdout, client.NewPageURL(results[idx].Link))

		return exitCodeOK, nil
	}

	pkg := pkgName(results[idx].Link)

	if opts.GoGet {
		cmd, args := buildGoGetCommand(pkg, "", false)
		if err := execCommand(cmd, args); err != nil {
			return exitCodeErrGoGet, fmt.Errorf("failed to get package: %w", err)
		}
	}

	fmt.Fprintln(os.Stdout, pkg)

	return exitCodeOK, nil
}

func pkgName(link string) string {
	return strings.Replace(link, "/", "", 1)
}

func buildGoGetCommand(pkg string, version string, update bool) (string, []string) {
	if version == "" {
		version = "latest"
	}

	pkgWithVer := fmt.Sprintf("%s@%s", pkg, version)

	if update {
		return "go", []string{"get", "-v", "-u", pkgWithVer}
	}

	return "go", []string{"get", "-v", pkgWithVer}
}

func execCommand(command string, args []string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command execution error: %w", err)
	}

	return nil
}
