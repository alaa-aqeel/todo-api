package config

import "fmt"

var App = Config{
	"APP_DEBUG": "false",
	"APP_ENV":   "production",
	"APP_URL":   "http://localhost:8080",

	// DATABASE
	"DB_USERNAME": "root s",
	"DB_PASSWORD": "",
	"DB_NAME":     "database",
	"DB_HOST":     "todo-api-db-dev-1",
	"DB_PORT":     "3306",
}

func (c Config) IsDebug() bool {

	return c.Get("APP_DEBUG") == "true"
}

func (c Config) GetDatabaseURL() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Baghdad",
		c.Get("DB_HOST"),
		c.Get("DB_USERNAME"),
		c.Get("DB_PASSWORD"),
		c.Get("DB_NAME"),
		c.Get("DB_PORT"),
	)
}
