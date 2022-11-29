package controllers

import (
	"backend/src/entities"
	enumhelper "backend/src/enum-helpers"
	httphelper "backend/src/http-helper"
	"backend/src/repositories"
	"backend/src/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UsersController struct {
	UsersRepository repositories.IUserRepository
}

func (this UsersController) CreateUser(writer http.ResponseWriter, request *http.Request) {
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

	if err = user.Prepare(); err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}
	user.Password = string(hashedPassword)

	userAlreadyExists, err := this.UsersRepository.GetUserByEmail(user.Email)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}
	if userAlreadyExists.Id != 0 {
		httphelper.HttpErrorResponse(writer, http.StatusConflict, errors.New(enumhelper.EmailHasBeenRegistered))
		return
	}

	insertedId, err := this.UsersRepository.Create(&user)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}

	httphelper.HttpResponse(writer, http.StatusCreated, map[string]int{"id": insertedId})
}

func (this UsersController) GetUser(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusBadRequest, err)
		return
	}

	user, err := this.UsersRepository.GetUserById(uint(userId))
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}

	if user.Id == 0 {
		httphelper.HttpErrorResponse(writer, http.StatusNotFound, errors.New(enumhelper.UserNotFound))
		return
	}

	user.Password = ""

	httphelper.HttpResponse(writer, http.StatusOK, user)
}
