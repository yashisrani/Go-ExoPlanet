package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/yashisrani/go-exoplanet/Handlers"
)

func main() {
	fmt.Println("app running on 8080")

	// create a new router using mux
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", handlers.AddExoPlanet).Methods("POST")
	r.HandleFunc("/allexoplanets", handlers.ListExoPlanet).Methods("GET")
	r.HandleFunc("/listexoplanets/{id}", handlers.GetExoPlanetByID).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handlers.UpdateExoPlanet).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
