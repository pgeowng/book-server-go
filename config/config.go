package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PgURL    string `envconfig:"PG_URL`
	HTTPAddr string `envconfig:"HTTP_ADDR"`
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
