package poppy

import (
	"encoding/json"
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
	Poppy{
		Id:          5,
		Name:        "eGolf",
		CarType:     "electric",
		ChargeLevel: 80,
		Latitude:    51.1986,
		Longitude:   4.3944,
	},
	Poppy{
		Id:          6,
		Name:        "eGolf",
		CarType:     "electric",
		ChargeLevel: 20,
		Latitude:    51.2117,
		Longitude:   4.4329,
	},
	Poppy{
		Id:          7,
		Name:        "A3 CNG",
		CarType:     "gaz",
		ChargeLevel: 60,
		Latitude:    51.1989,
		Longitude:   4.4008,
	},
	Poppy{
		Id:          8,
		Name:        "A3",
		CarType:     "gaz",
		ChargeLevel: 50,
		Latitude:    51.2139,
		Longitude:   4.3928,
	},
}

type OptionResponse struct {
	Leg   maas_api.Leg   `json:"leg"`
	Meta  Meta           `json:"meta"`
	Terms maas_api.Terms `json:"terms"`
}

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
	var res []OptionResponse
	// create stuct for the JSON
	for _, poppy := range poppiesMap {
		option := OptionResponse{
			Leg: maas_api.Leg{
				StartTime: time.Now().UTC(),
				EndTime:   time.Now().UTC(),
				From: maas_api.Coord{
					Lat: poppy.Latitude,
					Lon: poppy.Longitude,
				},
				To: maas_api.Coord{
					Lat: maasRequest.ToLat,
					Lon: maasRequest.ToLon,
				},
				Mode:     "CAR",
				AgencyID: "Poppy",
			},
			Meta: Meta{
				PoppyInfo: poppy,
			},
			Terms: maas_api.Terms{
				Price: maas_api.Price{
					Amount:   0.50,
					Currency: "EUR",
					Unit:     "MIN",
				},
			},
		}
		res = append(res, option)
	}
	// return JSON
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}

}
