package httpdoc

import (
	"net/http"
	"time"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// GetHTMLNode get html node
func GetHTMLNode(url string) (*html.Node, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	resp, err := (&http.Client{Timeout: 10 * time.Second}).Do(req)
	if err != nil {
		return nil, err
	}

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
