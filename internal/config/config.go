package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"githuh.com/cng-by-example/students/internal/db"
)

const (
	// Prefix indicates environment variables prefix.
	Prefix = "students_"
)

type Config struct {
	Listen   string    `koanf:"listen"`
	Database db.Config `koanf:"database"`
}

// New reads configuration with koanf.
func New() Config {
	var instance Config

	k := koanf.New(".")

	// load default configuration from its struct
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	// load environment variables
	// students_database__url
	// database__url
	// database.url

	// students_database__connection_timeout
	// database__connection_timeout
	// database.connection_timeout
	if err := k.Load(
		env.Provider(
			Prefix,
			".",
			func(s string) string {
				return strings.ReplaceAll(
					strings.ToLower(strings.TrimPrefix(s, Prefix)),
					"__", ".",
				)
			}), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	log.Printf("following configuration is loaded:\n%+v", instance)

	return instance
}
