package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status		string 		`json:"status"`
	Message		string		`json:"message"`
	Data		interface{}	`json:"data,omitempty"`
	Errors		interface{}	`json:"errors,omitempty"`
}

func HTTPResponse(w http.ResponseWriter, status string, statusCode int, message string, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Status: status,
        Message: message,
    }

	if data != nil {
		if status == "success" {
			response.Data = data
		}else {
			response.Errors = data
		}
	}

	json.NewEncoder(w).Encode(response)
}