package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"maas_webservice/maas_api"
	"maas_webservice/poppy"
	"net/http"
	"time"

	"github.com/gorilla/schema"
)

var AppImpl = App{}

type App struct {
}

func (a *App) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

type ResponseError struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type Foo struct {
	Bar string
}
type appAtoBRequest struct {
	FromLat float64 `schema:"fromLat"`
	FromLon float64 `schema:"fromLon"`
	ToLat   float64 `schema:"toLat"`
	ToLon   float64 `schema:"toLon"`
}
type aroundRequest struct {
	FromLat float64 `schema:"fromLat"`
	FromLon float64 `schema:"fromLon"`
}

func (a *App) getOptionsAround(w http.ResponseWriter, r *http.Request) {
	var res []maas_api.OptionResponse = getOptions(&poppy.TspPoppyImpl, maas_api.Coord{Lat: 52.32, Lon: 4.432})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
	// poppy.TspPoppyImpl.OptionsAround(maas_api.Coord{Lat: 52.32, Lon: 4.432})
}

func getOptions(m_func maas_api.MaasFunctions, pos maas_api.Coord) []maas_api.OptionResponse {
	return m_func.OptionsAround(pos)
}

func (a *App) getRouteAtoB(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	request := new(appAtoBRequest)
	if err := schema.NewDecoder().Decode(request, r.Form); err != nil {
		panic(err)
	}
	if request.FromLat == 0.0 || request.FromLon == 0.0 || request.ToLat == 0.0 || request.ToLon == 0.0 {
		resp := ResponseError{
			Error:   404,
			Message: "one parameter is missing",
		}
		log.Printf("%f, %f, %f, %f", request.FromLat, request.FromLon, request.ToLat, request.ToLon)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err)
		}
		return
	}

	url := fmt.Sprintf("http://connect2move-dev.azurewebsites.net/dm_mvp/staging/alpha/UserReq.php?token=dvgq5vnkk7py144fuq4s7x&startLat=%.4f&startLng=%.4f&endLat=%.4f&endLng=%.4f", request.FromLat, request.FromLon, request.ToLat, request.ToLon)
	responses := asyncHttpGets([]string{url})
	// jsonResponse := new(Foo)
	for _, resp := range responses {
		bodyBytes, _ := ioutil.ReadAll(resp.response.Body)
		bodyString := string(bodyBytes)
		resp.response.Body.Close()
		fmt.Fprint(w, bodyString)
	}
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered
	responses := []*HttpResponse{}
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}
