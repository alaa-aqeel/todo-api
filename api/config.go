package main

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	BaseDir string

	// App
	Name      string `envconfig:"APP_NAME" default:"todo-api"`
	Envirment string `envconfig:"APP_ENV" default:"prod"`
	Debug     bool   `envconfig:"APP_DEBUG" default:"false"`

	// Server
	HOST    string `envconfig:"APP_HOST" default:"0.0.0.0"`
	PORT    string `envconfig:"APP_PORT" default:"8080"`
	TimeOut int    `envconfig:"SERVER_TIME_OUT" default:"10"`

	DatabaseHost     string `envconfig:"DB_HOST" default:"localhost"`
	DatabasePort     int    `envconfig:"DB_PORT" default:"5432"`
	DatabaseName     string `envconfig:"DB_NAME" default:"app_db"`
	DatabaseUsername string `envconfig:"DB_USERNAME" default:"root"`
	DatabasePassword string `envconfig:"DB_PASSWORD" default:""`
}

func (cfg *Config) LoadConfig() {
	err := godotenv.Load(filepath.Join(cfg.BaseDir, ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file %s", err)
	}
	err = envconfig.Process("", cfg)
	if err != nil {
		log.Fatalln(err)
	}
}
