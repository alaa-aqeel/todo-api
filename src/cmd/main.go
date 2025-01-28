package main

import (
	"net/http"

	"github.com/alaa-aqeel/todo/src/api/http/router"
	"github.com/alaa-aqeel/todo/src/config"
	"github.com/alaa-aqeel/todo/src/database"
)

func main() {
	config.InitConfig()
	database.ConnectDatabase(config.App.GetDatabaseURL())
	database.MigrateModels(database.Orm)
	r := router.NewChiRouter()
	http.ListenAndServe(":8080", r)
}
