package ui

import (
	"fmt"

	"github.com/sheepla/gofind/client"
	"github.com/toqueteos/webbrowser"
)

func Open(link string) error {
	url := client.NewPageURL(link)
	if err := webbrowser.Open(url); err != nil {
		return fmt.Errorf("failed to open the URL %s: %w", url, err)
	}

	return nil
}
