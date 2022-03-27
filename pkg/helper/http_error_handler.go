package helper

import (
	"encoding/json"
	"net/http"
	"textgopher/models"
)

func HTTPErrorHandler(response http.ResponseWriter, body string, status int) {
	response.WriteHeader(status)
	errorResponse := models.HttpErrorResponse{
		Body: body,
	}
	jData, _ := json.Marshal(errorResponse)
	response.Write(jData)
}
