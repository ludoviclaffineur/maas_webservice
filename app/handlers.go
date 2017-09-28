package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var AppImpl = App{}

type App struct {
}

func (a *App) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

type Foo struct {
	Bar string
}

func (a *App) getRouteAtoB(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
	url := fmt.Sprintf("http://connect2move-dev.azurewebsites.net/dm_mvp/staging/alpha/UserReq.php?token=dvgq5vnkk7py144fuq4s7x&startLat=%.4f&startLng=%.4f&endLat=%.4f&endLng=%.4f", 51.1955, 4.3793, 51.1955, 4.4230)
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
