package controllers

import (
	"backend/src/database"
	enumhelper "backend/src/enum-helpers"
	httphelper "backend/src/http-helper"
	"backend/src/repositories"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func GetCryptos(writer http.ResponseWriter, request *http.Request) {
	var page uint
	urlQuery, err := strconv.ParseUint(request.URL.Query().Get("page"), 10, 64)
	if err != nil {
		page = 1
	}

	if urlQuery <= 1 {
		page = 1
	} else {
		page = uint(urlQuery)
	}

	db, err := database.Connect()
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, errors.New(enumhelper.ConnectDatabaseFailure))
		return
	}

	repository := repositories.NewCryptoRepository(db)

	cryptos, err := repository.GetPageOfCryptos(page)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}

	httphelper.HttpResponse(writer, http.StatusOK, cryptos)
}

func Vote(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var isLikeVoteAction bool = strings.HasSuffix(request.RequestURI, "up")

	cryptoId, err := strconv.ParseUint(params["cryptoId"], 10, 64)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusBadRequest, errors.New(enumhelper.InvalidParam))
		return
	}

	db, err := database.Connect()
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, errors.New(enumhelper.ConnectDatabaseFailure))
		return
	}

	repository := repositories.NewCryptoRepository(db)

	crypto, err := repository.GetCryptoById(uint(cryptoId))
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusNotFound, errors.New(enumhelper.CryptoNotFound))
		return
	}

	if !isLikeVoteAction && crypto.Votes >= 1 {
		crypto.Votes -= 1
	} else {
		crypto.Votes += 1
	}

	err = repository.Vote(uint(cryptoId), &crypto)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}

	httphelper.HttpResponse(writer, http.StatusNoContent, nil)
}
