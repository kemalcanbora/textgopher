package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"textgopher/models"
	"textgopher/pkg/helper"
	"textgopher/service"
)

func Upload(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var file models.File
	err := json.NewDecoder(request.Body).Decode(&file)
	if err != nil {
		helper.HTTPErrorHandler(response, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.S3I.UploadFile(awsClient, user.Email, file.Path)
	if err != nil {
		log.Fatal(err)
	}
	response.WriteHeader(http.StatusOK)
}
