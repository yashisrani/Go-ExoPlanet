package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	models "github.com/yashisrani/go-exoplanet/Models"
)

var store = models.NewExoPlanetStore()

// to give response of the error
func responsewitherror(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// validation code
func validationExoplanet(ExoPlanet *models.ExoPlanet) error {
	if ExoPlanet.Name == "" || ExoPlanet.Description == "" {
		return models.ErrInvalid
	}

	if ExoPlanet.Distance <= 10 || ExoPlanet.Distance >= 1000 {
		return models.ErrInvalid
	}

	if ExoPlanet.Radius <= 0.1 || ExoPlanet.Radius >= 10 {
		return models.ErrInvalid
	}

	if ExoPlanet.Type != models.GasGiant && ExoPlanet.Type != models.Terrestrial {
		return models.ErrInvalid
	}
	return nil

}

// Create API (Add exoplanet)
func AddExoPlanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet models.ExoPlanet

	// to decode json response which given by user
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid data"})
	}

	// to add validation for addExoplanet
	if err := validationExoplanet(&exoplanet); err != nil {
		responsewitherror(w, http.StatusBadRequest, "Invalid Data")
		return
	}

	// to generate unique id for exoplanet
	exoplanet.ID = uuid.NewString()
	store.ExoPlanets[exoplanet.ID] = exoplanet // to store exoplanet id in store

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanet)

}

// Get Api (list of all exoplanets)

func ListExoPlanet(w http.ResponseWriter, r *http.Request) {
	exoplanets := make([]models.ExoPlanet, 0, len(store.ExoPlanets))
	for _, plantes := range store.ExoPlanets {
		exoplanets = append(exoplanets, plantes)
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(exoplanets)
}

// get exoplante by id
func GetExoPlanetByID(w http.ResponseWriter, r *http.Request) {
	// to get id using mux.vars
	params := mux.Vars(r)
	id := params["id"]

	// to validate is id is present or not
	exoplanet, ok := store.ExoPlanets[id]
	if !ok {
		responsewitherror(w, http.StatusNotFound, models.ErrNotFound.Error())
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(exoplanet)
}

// update exoplanet by id
func UpdateExoPlanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// to get requested planet id, which we want to change data
	var reqPlanet models.ExoPlanet

	// to decode json request
	if err := json.NewDecoder(r.Body).Decode(&reqPlanet); err != nil {
		responsewitherror(w, http.StatusBadRequest, "Invalid request payload")
	}

	if err := validationExoplanet(&reqPlanet); err != nil {
		responsewitherror(w, http.StatusBadRequest, "Invalid exoplanet data")
	}

	// to validate is id is present or not
	_, ok := store.ExoPlanets[id]
	if !ok {
		responsewitherror(w, http.StatusNotFound, models.ErrNotFound.Error())
		return
	}

	reqPlanet.ID = id
	store.ExoPlanets[id] = reqPlanet
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(reqPlanet)

}

// delete api
func DeleteExoPlanet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// to validate is id is present or not
	p, ok := store.ExoPlanets[id]
	if !ok {
		responsewitherror(w, http.StatusNotFound, models.ErrNotFound.Error())
		return
	}

	// to delete
	delete(store.ExoPlanets, id)
	w.WriteHeader(http.StatusNoContent)

	// adding , which planet is removed
	json.NewEncoder(w).Encode(p)

}
