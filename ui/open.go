package ui

import (
	"fmt"
	"path"

	"github.com/toqueteos/webbrowser"
)

func Open(link string) error {
	url := path.Join("https://pkg.go.dev", link)
	if err := webbrowser.Open(url); err != nil {
		return fmt.Errorf("failed to open the URL %s: %w", url, err)
	}

	return nil
}
