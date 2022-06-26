package client_test

import (
	"testing"

	"github.com/sheepla/gofind/client"
)

//nolint:paralleltest
func TestNewURL(t *testing.T) {
	have := client.NewURL(&client.Param{
		Query:           "JSON",
		Limit:           20,
		SearchForSymbol: false,
	})
	want := "https://pkg.go.dev/search?limit=20&q=JSON"

	if have != want {
		t.Errorf("have=%s want=%s", have, want)
	}

	have = client.NewURL(&client.Param{
		Query:           "YAML",
		Limit:           20,
		SearchForSymbol: true,
	})
	want = "https://pkg.go.dev/search?limit=20&m=symbol&q=YAML"

	if have != want {
		t.Errorf("have=%s want=%s", have, want)
	}
}

//nolint:paralleltest
func TestSearch(t *testing.T) {
	results, err := client.Search(&client.Param{
		Query:           "JSON",
		Limit:           20,
		SearchForSymbol: false,
	})
	if err != nil {
		t.Errorf("An error occurred on Search func: %s", err)
	}

	for _, result := range results {
		if result.Name == "" {
			t.Error("name is empty")
		}

		if result.Link == "" {
			t.Error("link is empty")
		}

		// if result.Description == "" {
		// 	t.Error("description is empty")
		// }

		if result.License == "" {
			t.Error("license is empty")
		}

		if result.LicenseLink == "" {
			t.Error("link of license is empty")
		}
	}
}
