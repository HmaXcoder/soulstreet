package json

import (
	"encoding/json"
	"log"
	"net/http"
)

func SendJson(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func SendJsonError(w http.ResponseWriter, statusCode int, err error) {
	SendJson(w, statusCode, struct {
		Err string `json:"error"`
	} {
		Err: err.Error(),
	})
}