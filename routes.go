package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"PoppyOptions",
		"GET",
		"/poppy/bookings/options",
		PoppyOptions,
	},
	Route{
		"PoppyIndex",
		"GET",
		"/poppy/bookings",
		PoppyIndex,
	},
	Route{
		"PoppyCreate",
		"POST",
		"/poppy/bookings",
		PoppyCreate,
	},
	Route{
		"PoppyShow",
		"GET",
		"/poppy/bookings/{bookingId}",
		PoppyShow,
	},
	Route{
		"PoppyUpdate",
		"PUT",
		"/poppy/bookings/{bookingId}",
		PoppyUpdate,
	},
}
