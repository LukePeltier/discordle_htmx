package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lukepeltier/discordle/ent"
	"github.com/lukepeltier/discordle/internal/routes"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:discordle.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("Failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("entClient", client)
			return next(c)
		}
	})

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e.Use(middleware.Logger())

	routes.SetupRoutes(e)
	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":3000"))
}
