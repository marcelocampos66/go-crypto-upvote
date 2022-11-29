package controllers

import (
	"backend/src/entities"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type LoginMockUsersRepository struct{}

func (this LoginMockUsersRepository) Create(user *entities.User) (int, error) {
	return user.Id, nil
}

func (this LoginMockUsersRepository) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User

	user.Id = 1
	user.Name = "Unknow User"
	user.Email = "email@email.com"
	user.Password = "$2a$10$VPXIxMzPMiINuCc6IalYYegqjNOj1.y9pZ/KYXnFW5x0/RhiimMZW"

	return user, nil
}

func (this LoginMockUsersRepository) GetUserById(userId uint) (entities.User, error) {
	var user entities.User

	user.Id = 1
	user.Name = "Unknow User"
	user.Email = "email@email.com"

	return user, nil
}

func TestLoginController(t *testing.T) {
	t.Run("Login: Should return a token", func(t *testing.T) {
		stringfiedBody := `{"email":"email@email.com","password":"123456"}`
		request := httptest.NewRequest(http.MethodPost, "/v1/crypto-upvote/login", strings.NewReader(stringfiedBody))
		writer := httptest.NewRecorder()

		repository := LoginMockUsersRepository{}

		loginController := LoginController{
			UsersRepository: repository,
		}

		loginController.Login(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)
		body := map[string]string{"token": ""}
		json.Unmarshal(responseBody, &body)

		if !reflect.DeepEqual(http.StatusOK, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusOK, result.StatusCode)
		}
		if body["token"] == "" {
			t.Errorf("empty token: %s", body["token"])
		}
	})
}
