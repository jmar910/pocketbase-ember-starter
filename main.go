package main

import (
	"log"
	"os"
	"saturo-gojo/web"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	// uncomment once you have at least one .go migration file in the "migrations" directory
	// _ "yourpackage/migrations"
)

func main() {
	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		webHandler := web.NewHandler(e.App)
		webHandler.AddRoute(e.Router)

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
