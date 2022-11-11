package main

import (
	_ "gitlab.com/mCassy/gistclone/docs"
	"log"

	"gitlab.com/mCassy/gistclone/app/database"
	"gitlab.com/mCassy/gistclone/app/routes"
)

func main() {
	database.ConnectDb()
	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
