package main

import (
	"gistclone/app/database"
	"gistclone/app/routes"
	_ "gistclone/docs"
	"log"
)

func main() {
	database.ConnectDb()
	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
