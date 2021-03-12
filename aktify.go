package main

import (
	"fmt"
	"net/http"
	"projects/goaktify/handlers"
	"projects/goaktify/models"

	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	var c handlers.Campaign

	err := models.InitDB("postgres://user:pass@localhost/bookstore")
    if err != nil {
        log.Fatal(err)
    }

	c.Init()

	// Is our DB up? etc...
	// Would use this style to test handlers as well
	mux.HandleFunc(pat.Get("/healthcheck/"), c.HealthCheck)

	// CRUDL endpoints for our Campaigns

	// Create our campaign
	mux.HandleFunc(pat.Post("/campaigns/"), c.Create)

	// Read our campaign by ID
	mux.HandleFunc(pat.Get("/campaigns/:id"), c.Read)

	// Update our campaign.
	// Specifically no PATCH endpoint here.
	// Going to enforce idempotency for now
	mux.HandleFunc(pat.Put("/campaigns/"), c.Update)

	// Delete our campaign by ID
	mux.HandleFunc(pat.Delete("/campaigns/:id"), c.Delete)

	// List all campaigns. I would usally do this by company/location ID. Not going to here.
	mux.HandleFunc(pat.Get("/campaigns/"), c.List)

	http.ListenAndServe("localhost:8000", mux)
}





