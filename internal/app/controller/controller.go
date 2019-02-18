package controller

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseSuccess(w http.ResponseWriter, msg string, data interface{}) {
	ResponseJson(w, 200, "OK", msg, data)
}

func ResponseJson(w http.ResponseWriter, code int, status string, msg string, data interface{}) {
	resp := Response{
		Status:  status,
		Message: msg,
		Data:    data,
	}

	response, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		panic(jsonErr)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		panic(err)
	}
}
