package expand

import (
	"encoding/json"
	"net/http"

	"github.com/kcasas/short_url/internal/db"
	"github.com/kcasas/short_url/internal/db/models"
	"github.com/sirupsen/logrus"
)

type RequestPayload struct {
	Short string `json:"short"`
}

type JsonResponse struct {
	Long string `json:"long"`
}

func Handler(rw http.ResponseWriter, r *http.Request) {
	payload := RequestPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte(err.Error()))
		return
	}

	urlModel := models.URL{Short: payload.Short}
	err = urlModel.Expand(db.DB())

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(rw).Encode(&JsonResponse{
		Long: urlModel.LongURL,
	})
	logrus.Error(err)
}
