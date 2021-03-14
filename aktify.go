package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"projects/goaktify/handlers"
	"projects/goaktify/models"

	"github.com/joho/godotenv"

	"goji.io"
	"goji.io/pat"
)

func main() {
	err := initEnv()
	if err != nil {
		log.Fatal(err)
	}
	mux := goji.NewMux()
	var ch handlers.CampaignHandler

	err = models.InitDB(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	//TODO setup CORS

	// CRUDL endpoints for our Campaigns

	// Create campaign
	mux.HandleFunc(pat.Post("/campaigns/"), ch.Create)

	// Read campaign by ID
	mux.HandleFunc(pat.Get("/campaigns/:id"), ch.Read)

	// Update campaign
	// Specifically no PATCH endpoint here
	// Going to enforce idempotency for now by omitting
	mux.HandleFunc(pat.Put("/campaigns/"), ch.Update)

	// Delete our campaign by ID
	mux.HandleFunc(pat.Delete("/campaigns/:id"), ch.Delete)

	// List campaigns
	// I would usally do this by company or locationID
	// Not going to here due brevity
	mux.HandleFunc(pat.Get("/campaigns/"), ch.List)

	http.ListenAndServe(":8080", mux)
}
func initEnv() error {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	return err
}
