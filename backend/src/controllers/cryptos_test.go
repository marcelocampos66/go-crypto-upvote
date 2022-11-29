package controllers

import (
	"backend/src/entities"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type MockCryptoRepository struct {
	db *gorm.DB
}

func (this MockCryptoRepository) GetPageOfCryptos(page uint) ([]entities.Crypto, error) {
	cryptos := []entities.Crypto{
		{
			Id:    1,
			Name:  "Bitcoin",
			Code:  "BTC",
			Votes: 1,
		},
	}

	return cryptos, nil
}

func (this MockCryptoRepository) GetCryptoById(cryptoId uint) (entities.Crypto, error) {
	crypto := entities.Crypto{
		Id:    1,
		Name:  "Bitcoin",
		Code:  "BTC",
		Votes: 1,
	}

	return crypto, nil
}

func (this MockCryptoRepository) Vote(cryptoId uint, crypto *entities.Crypto) error {

	return nil
}

func TestCryptosController(t *testing.T) {
	t.Run("GetCryptos", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/crypto-upvote/cryptos", nil)
		writer := httptest.NewRecorder()

		repository := MockCryptoRepository{}

		cryptoController := CryptoController{
			CryptoRepository: repository,
		}

		cryptoController.GetCryptos(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)
		var body []entities.Crypto
		json.Unmarshal(responseBody, &body)

		if !reflect.DeepEqual(http.StatusOK, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusOK, result.StatusCode)
		}
		if !reflect.DeepEqual(len(body), 1) {
			t.Errorf("wanted: %d, got: %d", 1, len(body))
		}
		if body[0].Name != "Bitcoin" {
			t.Errorf("wanted: %s, got: %s", "Bitcoin", body[0].Name)
		}
		if body[0].Code != "BTC" {
			t.Errorf("wanted: %s, got: %s", "BTC", body[0].Code)
		}
	})

	t.Run("GetCrypto", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/crypto-upvote/cryptos/1", nil)
		writer := httptest.NewRecorder()

		vars := map[string]string{
			"cryptoId": "1",
		}

		request = mux.SetURLVars(request, vars)

		repository := MockCryptoRepository{}

		cryptoController := CryptoController{
			CryptoRepository: repository,
		}

		cryptoController.GetCrypto(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)
		var body entities.Crypto
		json.Unmarshal(responseBody, &body)

		if !reflect.DeepEqual(http.StatusOK, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusOK, result.StatusCode)
		}
		if body.Name != "Bitcoin" {
			t.Errorf("wanted: %s, got: %s", "Bitcoin", body.Name)
		}
		if body.Code != "BTC" {
			t.Errorf("wanted: %s, got: %s", "BTC", body.Code)
		}
	})

	t.Run("Vote Like", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPut, "/v1/crypto-upvote/cryptos/1/up", nil)
		writer := httptest.NewRecorder()

		vars := map[string]string{
			"cryptoId": "1",
		}

		request = mux.SetURLVars(request, vars)

		repository := MockCryptoRepository{}

		cryptoController := CryptoController{
			CryptoRepository: repository,
		}

		cryptoController.Vote(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)

		if !reflect.DeepEqual(http.StatusNoContent, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusNoContent, result.StatusCode)
		}
		if !reflect.DeepEqual("", string(responseBody)) {
			t.Errorf("wanted: %#v, got: %#v", "", string(responseBody))
		}
	})

	t.Run("Vote Deslike", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPut, "/v1/crypto-upvote/cryptos/1/down", nil)
		writer := httptest.NewRecorder()

		vars := map[string]string{
			"cryptoId": "1",
		}

		request = mux.SetURLVars(request, vars)

		repository := MockCryptoRepository{}

		cryptoController := CryptoController{
			CryptoRepository: repository,
		}

		cryptoController.Vote(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)

		if !reflect.DeepEqual(http.StatusNoContent, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusNoContent, result.StatusCode)
		}
		if !reflect.DeepEqual("", string(responseBody)) {
			t.Errorf("wanted: %#v, got: %#v", "", string(responseBody))
		}
	})
}
