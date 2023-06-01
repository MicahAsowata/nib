package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/pocketbase/dbx"
)

type Repo struct {
	app *fiber.App
	db  *dbx.DB
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	urlExample := "postgres://tsc:tsc-pswd@localhost:5432/tsc"
	db, _ := dbx.Open("postgres", urlExample)

	repo := &Repo{
		app: app,
		db:  db,
	}
	routes(repo)
	log.Println("Work done start")
	err := app.Listen("0.0.0.0:3030")
	if err != nil {
		log.Panic(err)
	}
}
