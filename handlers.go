package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var poppiesMap = []Poppy{
	Poppy{
		Id:          1,
		Name:        "eGolf",
		CarType:     "electric",
		ChargeLevel: 80,
		Latitude:    51.2085,
		Longitude:   4.4119,
	},
	Poppy{
		Id:          2,
		Name:        "eGolf",
		CarType:     "electric",
		ChargeLevel: 20,
		Latitude:    51.2069,
		Longitude:   4.4119,
	},
	Poppy{
		Id:          3,
		Name:        "A3 CNG",
		CarType:     "gaz",
		ChargeLevel: 60,
		Latitude:    51.2076,
		Longitude:   4.4151,
	},
	Poppy{
		Id:          4,
		Name:        "A3",
		CarType:     "gaz",
		ChargeLevel: 50,
		Latitude:    51.2104,
		Longitude:   4.4150,
	},
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func PoppyIndex(w http.ResponseWriter, r *http.Request) {

}

func PoppyUpdate(w http.ResponseWriter, r *http.Request) {

}

func PoppyShow(w http.ResponseWriter, r *http.Request) {

}

func PoppyCreate(w http.ResponseWriter, r *http.Request) {

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

type meta struct {
	PoppyInfo Poppy `json:"poppyInfo"`
}

type price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
type terms struct {
	Price price `json:"price"`
}

type optionResponse struct {
	Leg   Leg   `json:"leg"`
	Meta  meta  `json:"meta"`
	Terms terms `json:"terms"`
}

/*
PoppyOptions : to test : curl -i -H "Accept: application/json" "127.0.0.1:8080/poppy/bookings/options?fromLat=54.25"
*/
func PoppyOptions(w http.ResponseWriter, r *http.Request) {
	var maasRequest MaasRequest
	log.Printf("Je suis dedans")
	maasRequest.Mode = r.URL.Query().Get("mode")
	// Shit begin ---
	i, err := strconv.ParseFloat(r.URL.Query().Get("fromLat"), 64)
	maasRequest.FromLat = i
	maasRequest.FromLon, err = strconv.ParseFloat(r.URL.Query().Get("fromLon"), 64)
	maasRequest.ToLat, err = strconv.ParseFloat(r.URL.Query().Get("toLat"), 64)
	maasRequest.ToLon, err = strconv.ParseFloat(r.URL.Query().Get("toLong"), 64)
	// this should be done like this the parameters gathering
	if err != nil {
		// panic(err)
	}
	var res []optionResponse
	for _, poppy := range poppiesMap {
		option := optionResponse{
			Leg: Leg{
				StartTime: time.Now().UTC(),
				EndTime:   time.Now().UTC(),
				From: Coord{
					Lat: poppy.Latitude,
					Lon: poppy.Longitude,
				},
				To: Coord{
					Lat: maasRequest.ToLat,
					Lon: maasRequest.ToLon,
				},
				Mode:     "CAR",
				AgencyID: "Poppy",
			},
			Meta: meta{
				PoppyInfo: poppy,
			},
			Terms: terms{
				Price: price{
					Amount:   0.15,
					Currency: "EUR",
				},
			},
		}
		res = append(res, option)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
