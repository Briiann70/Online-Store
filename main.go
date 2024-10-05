package main

import (
	"test-online-store/config"
	"test-online-store/database"
	"test-online-store/routes"
)

func main() {
	config.LoadEnv()
	database.ConnectDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
