package poppy

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"maas_webservice/maas_api"

	"github.com/gorilla/schema"
)

type TspPoppy struct {
	maas_api.Tsp
}

var TspPoppyImpl = TspPoppy{maas_api.Tsp{Name: "poppy"}}

type Meta struct {
	PoppyInfo Poppy `json:"poppyInfo"`
}

func HttpAuthenticate(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("starting")
		fn(w, r)
		log.Println("completed")
	}
}

// type OptionResponse struct {
// 	maas_api.OptionResponse
// 	Meta Meta `json:"meta"`
// }

/*
PoppyOptions : to test : curl -i -H "Accept: application/json" "127.0.0.1:8080/poppy/bookings/options?fromLat=54.25"
*/
func (t *TspPoppy) OptionsList(w http.ResponseWriter, r *http.Request) {
	// gather URL param to struct
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	maasRequest := new(maas_api.MaasRequest)
	log.Printf("HERE")
	if err := schema.NewDecoder().Decode(maasRequest, r.Form); err != nil {
		panic(err)
	}
	var res []maas_api.OptionResponse = t.OptionsAround(maas_api.Coord{Lat: maasRequest.FromLat, Lon: maasRequest.FromLon})
	// create stuct for the JSON

	// return JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
func (t *TspPoppy) OptionsAround(position maas_api.Coord) []maas_api.OptionResponse {
	var res []maas_api.OptionResponse
	for _, poppy := range poppiesMap {
		agenceID := fmt.Sprintf("%d", poppy.Id)
		poppyJSON, _ := json.Marshal(poppy)
		option := maas_api.OptionResponse{
			Leg: maas_api.Leg{
				StartTime: time.Now().UTC(),
				EndTime:   time.Now().UTC(),
				From: maas_api.Coord{
					Lat: poppy.Latitude,
					Lon: poppy.Longitude,
				},
				To: maas_api.Coord{
					Lat: poppy.Latitude,
					Lon: poppy.Longitude,
				},
				Mode:     "car",
				AgencyID: agenceID,
			},

			Terms: maas_api.Terms{
				Price: maas_api.Price{
					Amount:   0.50,
					Currency: "EUR",
					Unit:     "MIN",
				},
			},
			Meta: string(poppyJSON),
		}
		res = append(res, option)
	}
	return res
}
