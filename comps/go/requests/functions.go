package requests

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// ProcessResponse Helper function to return a http response, performing some common checking
func ProcessResponse(w http.ResponseWriter, resp interface{}, ctx context.Context, err error) {
	if err != nil {
		log.WithError(err).Errorf("error in response")
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	jsonResp, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		log.WithError(err).Errorf("error while try to marhsal into json response")
	}
	w.Write(jsonResp)
}
