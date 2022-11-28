package middlewares

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type requestBasicLog struct {
	method string
	host   string
	url    string
}

type requestParamLog struct {
	requestQuery  interface{}
	requestParams interface{}
	requestBody   interface{}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var requestBasicLog requestBasicLog
		var requestParamsLog requestParamLog

		queryParam := request.URL.Query()
		requestParams := mux.Vars(request)
		requestBody := request.Body

		requestBasicLog.method = request.Method
		requestBasicLog.host = request.Host
		requestBasicLog.url = request.RequestURI

		requestParamsLog.requestQuery = queryParam
		requestParamsLog.requestParams = requestParams
		requestParamsLog.requestBody = requestBody

		log.Print(requestBasicLog, requestParamsLog)
		next(writer, request)
	}
}
