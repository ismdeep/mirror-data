package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/internal/store"
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

// GetDownloadLinks get download links
func GetDownloadLinks() ([]string, error) {
	doc, err := GetHTMLNode("https://go.dev/dl/")
	if err != nil {
		return nil, err
	}

	var results []string
	lst := htmlquery.Find(doc, `//table[@class="downloadtable"]//tr/td/a[@class="download"]/@href`)
	for _, node := range lst {
		link := htmlquery.InnerText(node)
		results = append(results, fmt.Sprintf("https://go.dev%v", link))
	}

	return results, nil
}

// GetVersion get version
func GetVersion(s string) string {
	s = strings.TrimSpace(s)
	v := strings.Split(s, ".")
	items := []string{
		v[0],
	}
	for _, s := range v[1:] {
		if '0' <= s[0] && s[0] <= '9' {
			items = append(items, s)
			continue
		}
		break
	}

	return strings.Join(items, ".")
}

func main() {
	storage := store.New("go", conf.Config.StorageCoroutineSize)

	originLinks, err := GetDownloadLinks()
	if err != nil {
		panic(err)
	}
	for _, originLink := range originLinks {
		if !strings.Contains(originLink, "/dl/") {
			continue
		}
		fileName := originLink[strings.Index(originLink, "/dl/")+4:]
		version := GetVersion(fileName)
		storage.Add(fmt.Sprintf("all/%v", fileName), originLink)
		storage.Add(fmt.Sprintf("dist/%v/%v", version, fileName), originLink)
	}

	storage.CloseAndWait()
}
