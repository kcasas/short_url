package shorten

import (
	"encoding/json"
	"net/http"

	"github.com/kcasas/short_url/internal/db"
	"github.com/kcasas/short_url/internal/urlconv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type JsonResponse struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

func Handler(rw http.ResponseWriter, r *http.Request) {
	payload := RequestPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte(err.Error()))
		return
	}

	err = payload.validate()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte(err.Error()))
		return
	}

	dbAdapter := db.NewAdapter(db.DB())
	prefixer := urlconv.NewRandomizer(
		viper.GetInt64("PREFIX_MIN"),
		viper.GetInt64("PREFIX_MAX"),
	)

	shortener := urlconv.NewShortener(dbAdapter, prefixer)

	shorturl, err := shortener.Shorten(payload.URL, payload.Expiration)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(rw).Encode(&JsonResponse{
		Short: shorturl,
		Long:  payload.URL,
	})
	if err != nil {
		logrus.Error(err)
	}
}
