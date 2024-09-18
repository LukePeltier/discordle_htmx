package main

import (
	"log"
	"os"

	"github.com/lukepeltier/discordle/internal/routes"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		routes.SetupRoutes(e.Router)
		e.Router.GET("/static/*", apis.StaticDirectoryHandler(os.DirFS("static/"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
