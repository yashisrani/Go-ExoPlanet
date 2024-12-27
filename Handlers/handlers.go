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
func responsewitherror(w http.ResponseWriter,code int,message string)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error":message})
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
	if err:= validationExoplanet(&exoplanet);err!=nil {
		responsewitherror(w,http.StatusBadRequest,"Invalid Data")
		return
	}

	// to generate unique id for exoplanet
	exoplanet.ID = uuid.NewString()
	store.ExoPlanets[exoplanet.ID] = exoplanet // to store exoplanet id in store

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanet)

}

// Get Api (list of all exoplanets)

func ListExoPlanet(w http.ResponseWriter, r *http.Request)  {
	exoplanets:= make([]models.ExoPlanet, 0 , len(store.ExoPlanets))
	for _, plantes := range store.ExoPlanets {
        exoplanets = append(exoplanets, plantes)
    }
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(exoplanets)
}


// get exoplante by id
func GetExoPlanetByID(w http.ResponseWriter, r *http.Request)  {
	// to validate id
	params:=mux.Vars(r)
	id:=params["id"]

	exoplanet,ok:=store.ExoPlanets[id]
	if !ok {
		responsewitherror(w,http.StatusNotFound,models.ErrNotFound.Error())
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(exoplanet)
}