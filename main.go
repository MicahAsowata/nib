package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pocketbase/dbx"
)

type Repo struct {
	app *fiber.App
	db  *dbx.DB
}

func main() {
	_ = godotenv.Load()
	app := fiber.New()
	app.Use(logger.New())
	dsn := os.Getenv("DSN")
	db, _ := dbx.Open("postgres", dsn)

	repo := &Repo{
		app: app,
		db:  db,
	}
	routes(repo)
	log.Println("Work done start")
	err := app.Listen(":3030")
	if err != nil {
		log.Panic(err)
	}
}
