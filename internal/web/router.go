package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kcasas/short_url/internal/db"
	"github.com/kcasas/short_url/internal/urlconv"
	"github.com/kcasas/short_url/internal/web/api/expand"
	"github.com/kcasas/short_url/internal/web/api/shorten"
	"github.com/spf13/viper"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/z/ping", func(rw http.ResponseWriter, _ *http.Request) {
		_, _ = rw.Write([]byte("ok"))
	})

	dbAdapter := db.NewAdapter(db.DB())

	prefixer := urlconv.NewRandomizer(
		viper.GetInt64("PREFIX_MIN"),
		viper.GetInt64("PREFIX_MAX"),
	)

	api := router.PathPrefix("/api").Subrouter()
	api.Handle(
		"/shorten",
		shorten.NewShortenHandler(dbAdapter, prefixer),
	).Methods(http.MethodPost)

	api.Handle(
		"/expand",
		expand.NewExpandHandler(dbAdapter),
	).Methods(http.MethodPost)

	return router
}
