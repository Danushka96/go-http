package handlers

import (
	"encoding/json"
	"net/http"
)

type GetAll struct {
}

func (GetAll) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response,_ := json.Marshal(PersonMap)
	_, _ = w.Write(response)
}