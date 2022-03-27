package routers

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"textgopher/models"
	"textgopher/pkg"
	"textgopher/pkg/helper"
	"textgopher/service"
	"time"
)

var user models.User

func SignUp(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(request.Body).Decode(&user)
	validate := validator.New()
	errVal := validate.Struct(user)
	if errVal != nil {
		helper.HTTPErrorHandler(response, "Email or Password field is empty!", http.StatusUnauthorized)
		return
	}

	userCheck, _ := service.User.FindWithEmail(mongoClient, user.Email)

	if userCheck.Email == user.Email {
		helper.HTTPErrorHandler(response, "This email is already registered!", http.StatusNotAcceptable)
		return
	}

	if err != nil {
		helper.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	user.Password = pkg.GetHash([]byte(user.Password))
	user.CreatedTime = time.Now().Unix()
	err = service.S3I.CreateBucket(awsClient, user.Email)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := service.User.Insert(mongoClient, user, "users")
	if err != nil {
		helper.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	json.NewEncoder(response).Encode(result)
}

func SignIn(response http.ResponseWriter, request *http.Request) {
	json.NewDecoder(request.Body).Decode(&user)
	response.Header().Set("Content-Type", "application/json")
	result, err := service.User.FindWithEmail(mongoClient, user.Email)
	if err != nil {
		helper.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	passErr := pkg.CheckPasswordHash(user.Password, result.Password)

	if passErr != true {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}

	jwtToken, err := pkg.GenerateJWT(result)
	if err != nil {
		helper.HTTPErrorHandler(response, `{"message":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))
}
