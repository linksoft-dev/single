package rest

import (
	"encoding/json"
	"net/http"
)

// GetRequestParams helper function to decode http params into object instance
func GetRequestParams(r *http.Request, dest interface{}) (err error) {
	err = json.NewDecoder(r.Body).Decode(dest)
	return
}

//Response helper function to wrapper http response
func Response(w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}
	if resp != nil {
		jsonResp, jsonErr := json.Marshal(resp)
		if jsonErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(jsonResp)
	}
	w.Write([]byte("ok"))
}
