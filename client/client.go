package client

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Param struct {
	Query           string
	Limit           int
	SearchForSymbol bool
}

type Result struct {
	Name        string `json:"name"`
	Link        string `json:"link"`
	License     string `json:"license"`
	LicenseLink string `json:"licenseceLink"`
	Description string `json:"description"`
	Snippet     string `json:"snippet"`
}

// nolint:varnamelen,exhaustruct,exhaustivestruct
func NewURL(param *Param) string {
	u := &url.URL{
		Scheme: "https",
		Host:   "pkg.go.dev",
		Path:   "search",
	}

	q := u.Query()
	q.Set("q", param.Query)
	q.Set("limit", strconv.Itoa(param.Limit))

	if param.SearchForSymbol {
		q.Set("m", "symbol")
	}

	u.RawQuery = q.Encode()

	return u.String()
}

func Search(param *Param) ([]Result, error) {
	res, err := http.Get(NewURL(param))
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the content: %w", err)
	}

	var (
		result  Result
		results []Result
	)

	doc.Find("div.SearchSnippet").Each(func(i int, div *goquery.Selection) {
		div.Find("h2").Each(func(i int, h2 *goquery.Selection) {
			result.Name = clean(h2.Text())
			result.Link = clean(h2.Find("a").AttrOr("href", ""))
		})

		result.Description = clean(div.Find("p[data-test-id=snippet-synopsis]").Text())

		license := div.Find("span[data-test-id=snippet-license] a")
		result.LicenseLink = clean(license.AttrOr("href", ""))
		result.License = clean(license.Text())

		div.Find("div.SearchSnippet-infoLabel").Each(func(i int, label *goquery.Selection) {
			result.Snippet = clean(label.Text())
		})

		results = append(results, result)
	})

	return results, nil
}

func clean(str string) string {
	s := strings.ReplaceAll(str, "\n", "")
	re := regexp.MustCompile(` +`)
	s = re.ReplaceAllString(s, " ")

	return strings.TrimSpace(s)
}

func NewPageURL(link string) string {
	return path.Join("https://pkg.go.dev", link)
}
