package main

import (
	"errors"

	"github.com/ismdeep/mirror-data/app/mirror/api"
)

func main() {
	if err := api.Run(); err != nil {
		panic(
			errors.Join(
				errors.New("api server failed"),
				err))
	}
}
