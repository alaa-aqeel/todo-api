package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alaa-aqeel/todo/src/helpers"
	"github.com/joho/godotenv"
)

type Config map[string]interface{}

func InitConfig() error {
	err := loadEnv(".env")
	if err != nil {
		return err
	}

	return nil
}

func (m Config) Get(key string) interface{} {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	value, exists := m[key]
	if exists {
		return ""
	}

	return value
}

func loadEnv(filename string) error {
	baseDir := helpers.BaseDir(".")
	file := filepath.Join(baseDir, filename)
	err := godotenv.Load(filepath.Join(baseDir, filename))
	if err != nil {
		return fmt.Errorf("[Error]: file %s does not exist", file)
	}

	return nil
}
