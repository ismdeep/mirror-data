package conf

import (
	"context"
	"errors"
	"os"

	"github.com/sethvargo/go-envconfig"
)

type config struct {
	Server struct {
		Bind string `env:"MIRROR_SERVER_BIND,default=0.0.0.0:9000"`
		Mode string `env:"MIRROR_SERVER_MODE,default=release"`
		Site string `env:"MIRROR_SERVER_SITE"`
	}
	System struct {
		Data          string `env:"MIRROR_SYSTEM_DATA,default=$HOME/Documents/mirror-data"`
		DownloadToken string `env:"MIRROR_DOWNLOAD_TOKEN,default="`
	}
	Notification struct {
		Relay struct {
			Endpoint string `env:"NOTIFICATION_RELAY_ENDPOINT,default=https://notification-relay.ismdeep.com"`
			Auth     string `env:"NOTIFICATION_RELAY_AUTH"`
		}
	}
}

// ROOT instance
var ROOT config

func init() {
	if err := envconfig.Process(context.TODO(), &ROOT); err != nil {
		panic(
			errors.Join(
				errors.New("failed to load config"),
				err))
	}

	if err := os.MkdirAll(ROOT.System.Data, 0750); err != nil {
		panic(
			errors.Join(
				errors.New("failed to create data folder"),
				err))
	}
}
