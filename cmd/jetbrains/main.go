package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ismdeep/mirror-data/internal"
)

var storage *internal.Storage

func init() {
	tmp, err := internal.NewStorage("jetbrains", 32)
	if err != nil {
		return
	}

	storage = tmp
	storage.StartConsumer()
}

type Product struct {
	Name     string           `json:"name"`
	Link     string           `json:"link"`
	Releases []ProductRelease `json:"releases"`
}

type ProductRelease struct {
	Build        string `json:"build"`
	Date         string `json:"date"`
	Downloads    map[string]ProductDownload
	MajorVersion string `json:"majorVersion"`
	Type         string `json:"type"`
	Version      string `json:"version"`
}

type ProductDownload struct {
	ChecksumLink string `json:"checksumLink"`
	Link         string `json:"link"`
	Size         int64  `json:"size"`
}

func Fetch(code string, productName string) error {
	reqURL := fmt.Sprintf("https://data.services.jetbrains.com/products?code=%v", code)
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return err
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return err
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var products []Product
	if err := json.Unmarshal(raw, &products); err != nil {
		return err
	}

	for _, product := range products {
		for _, release := range product.Releases {
			if release.Type != "release" {
				continue
			}
			for arch, download := range release.Downloads {
				originLink := download.Link
				fileName := originLink[strings.LastIndex(originLink, "/")+1:]
				if len(fileName) <= 5 || fileName[len(fileName)-5:] == ".json" {
					continue
				}
				link := fmt.Sprintf("%v/%v/%v/%v/%v/%v", productName, release.Type, release.MajorVersion, release.Version, arch, fileName)
				storage.Add(link, originLink)
			}
		}
	}

	return nil
}

func main() {
	products := map[string]string{
		"CL":  "CLion",
		"DG":  "DataGrip",
		"DS":  "DataSpell",
		"GO":  "GoLand",
		"IIU": "IDEA",
		"PC":  "PyCharm",
		"PCC": "PyCharmCommunity",
		"PS":  "PhpStorm",
		"RD":  "Rider",
		"RM":  "RubyMine",
		"WS":  "WebStorm",
	}

	for code, productName := range products {
		if err := Fetch(code, productName); err != nil {
			panic(err)
		}
	}
	close(storage.C)

	storage.WG.Wait()
	fmt.Println("Done")
}
