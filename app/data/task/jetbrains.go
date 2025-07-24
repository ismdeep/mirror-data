package task

import (
	"fmt"
	"strings"

	"github.com/ismdeep/mirror-data/app/data/conf"
	"github.com/ismdeep/mirror-data/app/data/global"
	"github.com/ismdeep/mirror-data/app/data/internal/httputil"
	"github.com/ismdeep/mirror-data/app/data/internal/store"
)

type JetBrains struct {
	storage *store.Storage
}

func (receiver *JetBrains) GetBucketName() string {
	return "jetbrains"
}

func (receiver *JetBrains) Run() {
	receiver.storage = store.New("jetbrains", conf.Config.StorageCoroutineSize)

	defer func() {
		close(receiver.storage.C)
		receiver.storage.WG.Wait()
	}()

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
		if err := receiver.fetch(code, productName); err != nil {
			global.Errors <- err
			return
		}
	}
}
func (receiver *JetBrains) fetch(code string, productName string) error {
	var products []Product
	reqURL := fmt.Sprintf("https://data.services.jetbrains.com/products?code=%v", code)
	if err := httputil.GetWithJSONUnmarshal(reqURL, &products); err != nil {
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
				receiver.storage.Add(link, originLink)
			}
		}
	}

	return nil
}

// Product model
type Product struct {
	Name     string           `json:"name"`
	Link     string           `json:"link"`
	Releases []ProductRelease `json:"releases"`
}

// ProductRelease model
type ProductRelease struct {
	Build        string `json:"build"`
	Date         string `json:"date"`
	Downloads    map[string]ProductDownload
	MajorVersion string `json:"majorVersion"`
	Type         string `json:"type"`
	Version      string `json:"version"`
}

// ProductDownload download model
type ProductDownload struct {
	ChecksumLink string `json:"checksumLink"`
	Link         string `json:"link"`
	Size         int64  `json:"size"`
}
