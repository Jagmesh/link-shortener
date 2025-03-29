package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, status int, responseData any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(responseData)
}
