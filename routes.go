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
		"DocumentIndex",
		"GET",
		"/documents",
		DocumentIndex,
	},
	Route{
		"DocumentCreate",
		"POST",
		"/documents",
		DocumentCreate,
	},
	Route{
		"DocumentShow",
		"GET",
		"/documents/{documentId}",
		DocumentShow,
	},
}
