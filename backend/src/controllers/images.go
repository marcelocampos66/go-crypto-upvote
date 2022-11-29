package controllers

import (
	enumhelper "backend/src/enum-helpers"
	httphelper "backend/src/http-helper"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type cryptoCodesUrl struct {
	code string
	url  string
}

var cryptoUrls = []cryptoCodesUrl{
	{
		code: "BTC",
		url:  "/src/images/bitcoin_logo.jpg",
	},
	{
		code: "ETH",
		url:  "/src/images/ethereum_logo.jpg",
	},
	{
		code: "USDC",
		url:  "/src/images/usd_coin_logo.jpg",
	},
	{
		code: "XRP",
		url:  "/src/images/xrp_logo.jpg",
	},
	{
		code: "DOGE",
		url:  "/src/images/doge_logo.jpg",
	},
	{
		code: "ADA",
		url:  "/src/images/ada_logo.jpg",
	},
	{
		code: "LTC",
		url:  "/src/images/litecoin_logo.jpg",
	},
	{
		code: "BCH",
		url:  "/src/images/bitcoin_cash_logo.jpg",
	},
	{
		code: "SHIB",
		url:  "/src/images/shiba_inu_logo.jpg",
	},
	{
		code: "ATOM",
		url:  "/src/images/cosmos_logo.jpg",
	},
}

type ImagesController struct{}

func (this ImagesController) GetImage(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	cryptoCode := strings.ToUpper(params["cryptoCode"])
	if cryptoCode == "" {
		httphelper.HttpErrorResponse(writer, http.StatusBadRequest, errors.New(enumhelper.InvalidCryptoCode))
		return
	}

	var relativePath string
	for i := range cryptoUrls {
		if cryptoUrls[i].code == cryptoCode {
			relativePath = cryptoUrls[i].url
		}
	}

	if relativePath == "" {
		httphelper.HttpErrorResponse(writer, http.StatusBadRequest, errors.New(enumhelper.InvalidCryptoCode))
		return
	}

	rootPath, _ := os.Getwd()
	filePath := fmt.Sprintf("%s%s", rootPath, relativePath)
	log.Print("filePath", filePath)

	http.ServeFile(writer, request, filePath)
}
