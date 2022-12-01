package controllers

import (
	"backend/src/entities"
	enumhelper "backend/src/enum-helpers"
	httphelper "backend/src/http-helper"
	"backend/src/repositories"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

type CryptoController struct {
	CryptoRepository repositories.ICryptoRepository
}

func (this CryptoController) GetCryptos(writer http.ResponseWriter, request *http.Request) {
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

	cryptos, err := this.CryptoRepository.GetPageOfCryptos(page)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(cryptos))

	var updatedCryptos []entities.Crypto

	go func() {
		for _, crypto := range cryptos {
			updatedCrypto := this.CryptoRepository.GetCryptoLastQuotation(crypto, &waitGroup)
			updatedCryptos = append(updatedCryptos, updatedCrypto)
		}
	}()

	waitGroup.Wait()

	httphelper.HttpResponse(writer, http.StatusOK, updatedCryptos)
}

func (this CryptoController) GetCrypto(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	cryptoId, err := strconv.ParseUint(params["cryptoId"], 10, 64)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusBadRequest, errors.New(enumhelper.InvalidParam))
		return
	}

	crypto, err := this.CryptoRepository.GetCryptoById(uint(cryptoId))
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusNotFound, errors.New(enumhelper.CryptoNotFound))
		return
	}

	httphelper.HttpResponse(writer, http.StatusOK, crypto)
}

func (this CryptoController) Vote(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var isLikeVoteAction bool = strings.HasSuffix(request.RequestURI, "up")

	cryptoId, err := strconv.ParseUint(params["cryptoId"], 10, 64)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusBadRequest, errors.New(enumhelper.InvalidParam))
		return
	}

	crypto, err := this.CryptoRepository.GetCryptoById(uint(cryptoId))
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusNotFound, errors.New(enumhelper.CryptoNotFound))
		return
	}

	if !isLikeVoteAction && crypto.Votes >= 1 {
		crypto.Votes -= 1
	} else {
		crypto.Votes += 1
	}

	err = this.CryptoRepository.Vote(uint(cryptoId), &crypto)
	if err != nil {
		httphelper.HttpErrorResponse(writer, http.StatusInternalServerError, err)
		return
	}

	httphelper.HttpResponse(writer, http.StatusNoContent, nil)
}
