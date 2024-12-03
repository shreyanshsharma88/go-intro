package utils

import (
	"encoding/json"
	"net/http"
)

func SendJsonResponse(res http.ResponseWriter, status int, data interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(data)
}
