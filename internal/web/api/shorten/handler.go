package shorten

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/kcasas/short_url/internal/db"
	"github.com/kcasas/short_url/internal/urlconv"
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
		rw.Write([]byte(err.Error()))
		return
	}

	err = payload.validate()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	dbAdapter := db.NewAdapter(db.DB())
	prefixer := urlconv.NewRandomizer(
		int64(math.Pow(62, 2)),
		int64(math.Pow(62, 3)-1),
	)

	shortener := urlconv.NewShortener(dbAdapter, prefixer)

	shorturl, err := shortener.Shorten(payload.URL, payload.Expiration)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(&JsonResponse{
		Short: shorturl,
		Long:  payload.URL,
	})
}
