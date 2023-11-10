package openssl

import (
	"fmt"

	"github.com/ismdeep/mirror-data/conf"
	"github.com/ismdeep/mirror-data/internal/store"
)

func Run() error {
	storage := store.New("openssl", conf.Config.StorageCoroutineSize)
	versions := []string{
		"1.1.0a",
		"1.1.0b",
		"1.1.0c",
		"1.1.0d",
		"1.1.0e",
		"1.1.0f",
		"1.1.0g",
		"1.1.0h",
		"1.1.0i",
		"1.1.0j",
		"1.1.0k",
		"1.1.0l",
		"1.1.1a",
		"1.1.1b",
		"1.1.1c",
		"1.1.1d",
		"1.1.1e",
		"1.1.1f",
		"1.1.1g",
		"1.1.1h",
		"1.1.1i",
		"1.1.1j",
		"1.1.1k",
		"1.1.1l",
		"1.1.1m",
		"1.1.1n",
		"1.1.1o",
		"1.1.1p",
		"1.1.1q",
		"1.1.1s",
		"1.1.1t",
		"1.1.1u",
		"1.1.1v",
		"1.1.1w",
	}
	for _, version := range versions {
		storage.Add(
			fmt.Sprintf("openssl-%v.tar.gz", version),
			fmt.Sprintf("https://www.openssl.org/source/openssl-%v.tar.gz", version))
	}

	storage.CloseAndWait()

	return nil
}
