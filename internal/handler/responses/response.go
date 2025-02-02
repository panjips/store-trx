package responses

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status		string 		`json:"status"`
	Message		string		`json:"message"`
	Data		interface{}	`json:"data,omitempty"`
}

func HTTPResponse(w http.ResponseWriter, statusCode int, message string, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
        Status:  "error",
        Message: message,
    }

	if data != nil {
		response.Status = "success"
		response.Data = data
	}

	json.NewEncoder(w).Encode(response)
}