package main

import (
	"log"
	"net/http"
	"os"
	"projects/goaktify/handlers"
	"projects/goaktify/models"

	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	var ch handlers.CampaignHandler

	// err := models.InitDB("postgres://user:pass@localhost/bookstore")
	err := models.InitDB(os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	ch.Init()

	//TODO add CORS

	// Is our DB up? etc...
	// Would use this style to test handlers as well
	mux.HandleFunc(pat.Get("/healthcheck/"), ch.HealthCheck)

	// CRUDL endpoints for our Campaigns

	// Create our campaign
	mux.HandleFunc(pat.Post("/campaigns/"), ch.Create)

	// Read our campaign by ID
	mux.HandleFunc(pat.Get("/campaigns/:id"), ch.Read)

	// Update our campaign.
	// Specifically no PATCH endpoint here.
	// Going to enforce idempotency for now
	mux.HandleFunc(pat.Put("/campaigns/"), ch.Update)

	// Delete our campaign by ID
	mux.HandleFunc(pat.Delete("/campaigns/:id"), ch.Delete)

	// List all campaigns. I would usally do this by company/location ID. Not going to here.
	mux.HandleFunc(pat.Get("/campaigns/"), ch.List)

	http.ListenAndServe("localhost:8000", mux)
}
