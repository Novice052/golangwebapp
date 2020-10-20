package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(NotFound)
	// Begin OpenAPI specific routing
	//sh := http.StripPrefix("/help/", http.FileServer(http.Dir("./openapi")))
	//router.PathPrefix("/help/").Handler(sh)
	// End OpenAPI specific routing
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
