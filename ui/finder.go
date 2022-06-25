package ui

import (
	"fmt"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/mattn/go-runewidth"
	"github.com/sheepla/gofind/client"
)

func Find(results []client.Result) (int, error) {
	// nolint:wrapcheck
	return fuzzyfinder.Find(
		results,
		func(idx int) string {
			return results[idx].Name
		},
		fuzzyfinder.WithPreviewWindow(func(idx, width, height int) string {
			if idx == -1 {
				return ""
			}

			// nolint:gomnd
			wrapedWidth := width/2 - 5

			return runewidth.Wrap(renderPreviewWindow(&results[idx]), wrapedWidth)
		}),
	)
}

func renderPreviewWindow(result *client.Result) string {
	return fmt.Sprintf(
		"%s\n\n\n%s\n\n%s",
		result.Name,
		result.Description,
		result.Snippet,
	)
}
