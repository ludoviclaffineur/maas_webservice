package poppy

import (
	"maas_webservice/route"
)

func (t *TspPoppy) GetRoutes() route.Routes {
	return route.Routes{
		route.Route{
			Name:        t.Name + "Options",
			Method:      "GET",
			Pattern:     "/" + t.Name + "/bookings/options",
			HandlerFunc: t.OptionsList,
		},
		route.Route{
			Name:        t.Name + "Index",
			Method:      "GET",
			Pattern:     "/" + t.Name + "/bookings",
			HandlerFunc: HttpAuthenticate(t.Index),
		},
		route.Route{
			Name:        t.Name + "Create",
			Method:      "POST",
			Pattern:     "/" + t.Name + "/bookings",
			HandlerFunc: HttpAuthenticate(t.Create),
		},
		route.Route{
			Name:        t.Name + "Show",
			Method:      "GET",
			Pattern:     "/" + t.Name + "/bookings/{bookingId}",
			HandlerFunc: HttpAuthenticate(t.Show),
		},
		route.Route{
			Name:        t.Name + "Update",
			Method:      "PUT",
			Pattern:     "/" + t.Name + "/bookings/{bookingId}",
			HandlerFunc: HttpAuthenticate(t.Update),
		},
	}
}
