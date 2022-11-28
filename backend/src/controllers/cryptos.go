package controllers

import (
	"backend/src/database"
	httphelper "backend/src/http-helper"
	"backend/src/repositories"
	"net/http"
	"strconv"
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
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
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
