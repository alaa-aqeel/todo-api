package main

import (
	"github.com/alaa-aqeel/todo-api/core"
	"github.com/alaa-aqeel/todo-api/helpers"
)

func main() {
	app := core.NewApp(&core.Config{
		BaseDir: helpers.GetBaseDir("api"),
	})
	app.Config.LoadConfig(".env")
	app.FasthttpRunServer()
}
