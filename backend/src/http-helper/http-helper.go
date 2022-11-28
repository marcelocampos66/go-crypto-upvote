package httphelper

import (
	"encoding/json"
	"log"
	"net/http"
)

func HttpResponse(writer http.ResponseWriter, statusCode int, responseBody interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if responseBody == nil {
		return
	}

	if err := json.NewEncoder(writer).Encode(responseBody); err != nil {
		log.Fatal(err)
	}
}

func HttpErrorResponse(writer http.ResponseWriter, statusCode int, err error) {
	HttpResponse(writer, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
}
