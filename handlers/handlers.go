package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/goaktify/models"
	"strconv"

	"goji.io/pat"
)

type Handler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type CampaignHandler struct{}

func (ch *CampaignHandler) Create(w http.ResponseWriter, r *http.Request) {

	var in models.Campaign

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = in.Create()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)

	return
}

func (ch *CampaignHandler) Read(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(pat.Param(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var campaign models.Campaign

	err = campaign.Read(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(campaign)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func (ch *CampaignHandler) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Update endpoint unimplemented")
	return
}

func (ch *CampaignHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(pat.Param(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var campaign models.Campaign

	err = campaign.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO handle case where its already been deleted more gracefully
	w.WriteHeader(200)
	return

}

func (ch *CampaignHandler) List(w http.ResponseWriter, r *http.Request) {
	fmt.Print("List endpoint unimplemented since companies don't exist")
	return
}
