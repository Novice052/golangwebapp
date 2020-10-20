package main

import (
	"github.com/rs/cors"
	"net/http"
)

func main() {
	router := NewRouter()
	handler := cors.Default().Handler(router)
	http.ListenAndServe(":3000", handler)
}
