package main

import (
	"bookstore/ent"
	"bookstore/router"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New()
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=BookStoreManagement password=Password123#@! sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	router.SetupRoutes(app, client)
	app.Listen(":3000")
}
