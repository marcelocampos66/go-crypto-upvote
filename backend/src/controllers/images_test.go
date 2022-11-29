package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestImagesController(t *testing.T) {
	t.Run("GetImage: Should get a Image", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/crypto-upvote/images/BTC", nil)
		writer := httptest.NewRecorder()

		vars := map[string]string{
			"cryptoCode": "BTC",
		}

		request = mux.SetURLVars(request, vars)

		imagesController := ImagesController{}

		imagesController.GetImage(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		if !reflect.DeepEqual(http.StatusOK, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusOK, result.StatusCode)
		}
	})

	t.Run("GetImage: Should return 'Invalid crypto code' error", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/crypto-upvote/images/invalidUnexistingParam", nil)
		writer := httptest.NewRecorder()

		imagesController := ImagesController{}

		imagesController.GetImage(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		var body map[string]string
		responseBody, _ := ioutil.ReadAll(result.Body)
		json.Unmarshal(responseBody, &body)

		if !reflect.DeepEqual(http.StatusBadRequest, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusOK, result.StatusCode)
		}
		if !reflect.DeepEqual("Invalid crypto code", body["erro"]) {
			t.Errorf("wanted: %s, got: %s", "Invalid crypto code", body["erro"])
		}
	})
}
