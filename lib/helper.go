package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	r := ErrorResponse{Message: msg}
	j, _ := json.Marshal(r)
	fmt.Fprint(w, string(j))
}

func WriteSuccess(w http.ResponseWriter, msg interface{}) {
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(msg)
	if err != nil {
		WriteError(w, 500, "Error while marshalling response to json")
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, string(b))
}
