package app

import "maas_webservice/route"

func (t *App) GetRoutes() route.Routes {
	return route.Routes{
		route.Route{
			Name:        "/",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: t.index,
		},
		route.Route{
			Name:        "/getRouteAtoB",
			Method:      "GET",
			Pattern:     "/route",
			HandlerFunc: t.getRouteAtoB,
		},
		route.Route{
			Name:        "/getRouteAtoB",
			Method:      "GET",
			Pattern:     "/around",
			HandlerFunc: t.getOptionsAround,
		},
	}
}
