package main

import (
	"log"
	"net/http"

	"github.com/glennsills/yact/config"
	"github.com/glennsills/yact/db/dbUtilities"
	"github.com/glennsills/yact/ui"
)

func main() {
	config.InitAppConfig()
	ui.AssignHandlers()
	dbUtilities.MigrateDb()
	srv := &http.Server{
		Addr:    config.App.Addr,
		Handler: ui.AssignHandlers(),
	}

	log.Fatal(srv.ListenAndServe())
}
