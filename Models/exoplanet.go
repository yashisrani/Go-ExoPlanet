package models

import (
	"errors"
)

type ExoPlanetType string

const (
	GasGiant    ExoPlanetType = "GasGiant"
	Terrestrial ExoPlanetType = "Terrestrial"
)

type ExoPlanet struct {
	ID          string        `json:"id,omitempty"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    float64       `json:"distance"`
	Radius      float64       `json:"radius"`
	Type        ExoPlanetType `json:"type"`
	Mass        *float64      `json:"mass,omitempty"`
}

// to store all exoplanets information in form of map (locally)
type ExoPlanetStore struct {
	ExoPlanets map[string]ExoPlanet
}

// to give response of added exoplanet information
func NewExoPlanetStore() *ExoPlanetStore {
	return &ExoPlanetStore{
		ExoPlanets: make(map[string]ExoPlanet),
	}
}

// for error handling
var (
	ErrNotFound = errors.New("exoplanet not found")
	ErrInvalid  = errors.New("invalid exoplanet")
)
