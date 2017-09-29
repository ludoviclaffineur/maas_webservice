package maas_api

import (
	"log"
	"net/http"
	"time"

	"maas_webservice/route"
)

type Tsp struct {
	Name string
}

func HttpAuthenticate(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("starting")
		fn(w, r)
		log.Println("completed")
	}
}

type MaasApi interface {
	Index(w http.ResponseWriter, r *http.Request)
	OptionsList(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetRoutes() route.Routes
}

type OptionResponse struct {
	Leg   Leg    `json:"leg"`
	Terms Terms  `json:"terms"`
	Meta  string `json:"meta"`
}

type MaasFunctions interface {
	// OptionsAtoB(from Coord, to Coord)
	OptionsAround(position Coord) []OptionResponse
}

type MaasRequest struct {
	Mode      string    `schema:"mode"`
	FromLat   float64   `schema:"fromLat"`
	FromLon   float64   `schema:"fromLon"`
	ToLat     float64   `schema:"toLat"`
	ToLon     float64   `schema:"toLon"`
	StartTime time.Time `schema:"startTime"`
	EndTime   time.Time `schema:"endTime"`
}
type Leg struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	From      Coord     `json:"from"`
	To        Coord     `json:"to"`
	Mode      string    `json:"mode"`
	AgencyID  string    `json:"agencyId"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Meta struct {
}

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Unit     string  `json:"unit"`
}

type Terms struct {
	Price Price `json:"price"`
}

func (t *Tsp) Index(w http.ResponseWriter, r *http.Request) {
}

func (t *Tsp) OptionsList(w http.ResponseWriter, r *http.Request) {
}

func (t *Tsp) Create(w http.ResponseWriter, r *http.Request) {
}

func (t *Tsp) Show(w http.ResponseWriter, r *http.Request) {
}

func (t *Tsp) Update(w http.ResponseWriter, r *http.Request) {
}
