package expand

import (
	"encoding/json"
	"net/http"

	"github.com/kcasas/short_url/internal/db"
	"github.com/sirupsen/logrus"
)

type RequestPayload struct {
	Short string `json:"short"`
}

type JsonResponse struct {
	Long string `json:"long"`
}

type ExpandHandler struct {
	db db.DBAdapter
}

func NewExpandHandler(adapter db.DBAdapter) *ExpandHandler {
	return &ExpandHandler{adapter}
}

func (handler *ExpandHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	payload := RequestPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		logrus.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte(err.Error()))
		return
	}

	longURL, err := handler.db.ExpandURL(payload.Short)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(rw).Encode(&JsonResponse{
		Long: longURL,
	})

	log := logrus.WithFields(logrus.Fields{
		"short": payload.Short,
		"long":  longURL,
	})
	log.Debug("/api/expand")

	if err != nil {
		log.Error(err)
	}
}
