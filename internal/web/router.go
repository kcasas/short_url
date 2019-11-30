package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcasas/short_url/internal/web/api/expand"
	"github.com/kcasas/short_url/internal/web/api/shorten"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/z/ping", func(rw http.ResponseWriter, _ *http.Request) {
		rw.Write([]byte("ok"))
	})

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/shorten", shorten.Handler).Methods(http.MethodPost)
	api.HandleFunc("/expand", expand.Handler).Methods(http.MethodPost)

	return router
}
