package controllers

import (
	"backend/src/authentication"
	"backend/src/entities"
	enumhelper "backend/src/enum-helpers"
	httphelper "backend/src/http-helper"
	"backend/src/repositories"
	"backend/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type LoginController struct {
	UsersRepository repositories.IUserRepository
}

func (this LoginController) Login(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusUnprocessableEntity, err)
		return
	}

	var user entities.User
	if err = json.Unmarshal(body, &user); err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusBadRequest, err)
		return
	}

	userOnDb, err := this.UsersRepository.GetUserByEmail(user.Email)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userOnDb.Password, user.Password); err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusUnauthorized, errors.New(enumhelper.EmailOrPasswordWrong))
		return
	}

	token, err := authentication.CreateToken(uint(userOnDb.Id))
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}

	httphelper.HttpResponse(writer, http.StatusOK, map[string]string{"token": token})
}
