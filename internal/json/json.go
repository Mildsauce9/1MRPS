package json

import (
	"encoding/json"
	"net/http"
)

// this is the writer function to start sending out response.
func Write(w http.ResponseWriter, status int, data any) {
	//set the application type so we return json response
	w.Header().Set("Content-Type", "application/json")
	//set the response code
	w.WriteHeader(http.StatusOK)
	//return the response
	json.NewEncoder(w).Encode(data)
}
