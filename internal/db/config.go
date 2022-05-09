package db

import "time"

type Config struct {
	URL               string        `koanf:"url"`
	ConnectionTimeout time.Duration `koanf:"connection_timeout"`
}
