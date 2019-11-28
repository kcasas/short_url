package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/z/ping", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write([]byte("ok"))
	})

	api := router.PathPrefix("/api").Subrouter()
	api.Methods(http.MethodPost).Headers("Content-Type", "application/json")
	api.HandleFunc("/shorten", ShortenHandler)

	return router
}
