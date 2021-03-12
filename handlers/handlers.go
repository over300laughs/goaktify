package handlers

import (
	"fmt"
	"net/http"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	HealthChecker(w http.ResponseWriter, r *http.Request)
}

type Campaign struct{}

func (c *Campaign) Init() {
	// nothing to initialize for now so just return
	return
}

func (c *Campaign) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Create endpoint")
	return
}

func (c *Campaign) Read(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Read endpoint")
	return
}

func (c *Campaign) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Update endpoint")
	return
}

func (c *Campaign) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Delete endpoint")
	return
}

func (c *Campaign) List(w http.ResponseWriter, r *http.Request) {
	fmt.Print("List endpoint\n")
	return
}

func (c *Campaign) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Print("List endpoint\n")
	return
}
