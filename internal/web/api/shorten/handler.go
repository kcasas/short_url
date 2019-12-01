package shorten

import (
	"encoding/json"
	"net/http"

	"github.com/kcasas/short_url/internal/urlconv"
	"github.com/sirupsen/logrus"
)

type JsonResponse struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

type ShortenHandler struct {
	db       urlconv.DBAdapter
	prefixer urlconv.Prefixer
}

func NewShortenHandler(adapter urlconv.DBAdapter, prefixer urlconv.Prefixer) *ShortenHandler {
	return &ShortenHandler{adapter, prefixer}
}

func (handler *ShortenHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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

	shortener := urlconv.NewShortener(handler.db, handler.prefixer)
	shorturl, err := shortener.Shorten(payload.URL, payload.Expiration)
	if err != nil {
		logrus.Error(err)
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
