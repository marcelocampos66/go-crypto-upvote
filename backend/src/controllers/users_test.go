package controllers

import (
	"backend/src/entities"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type MockUsersRepository struct{}

func (this MockUsersRepository) Create(user *entities.User) (int, error) {

	user.Id = 1

	return user.Id, nil
}

func (this MockUsersRepository) GetUserByEmail(email string) (entities.User, error) {
	var user entities.User

	return user, nil
}

func (this MockUsersRepository) GetUserById(userId uint) (entities.User, error) {
	var user entities.User

	user.Id = 1
	user.Name = "Random User"
	user.Email = "email@email.com"

	return user, nil
}

func TestUsersController(t *testing.T) {
	t.Run("CreateUser: Should create a user", func(t *testing.T) {
		stringfiedBody := `{"name":"Unknow User","email":"email@email.com","password":"123456"}`
		request := httptest.NewRequest(http.MethodGet, "/v1/crypto-upvote/users", strings.NewReader(stringfiedBody))
		writer := httptest.NewRecorder()

		repository := MockUsersRepository{}

		usersController := UsersController{
			UsersRepository: repository,
		}

		usersController.CreateUser(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)
		body := struct{ ID uint }{}
		json.Unmarshal(responseBody, &body)

		if !reflect.DeepEqual(http.StatusCreated, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusCreated, result.StatusCode)
		}
		if body.ID != 1 {
			t.Errorf("wanted: %d, got: %d", 1, body.ID)
		}
	})

	t.Run("CreateUser: Should return a error if dont pass email to register", func(t *testing.T) {
		stringfiedBody := `{"name":"Unknow User","password":"123456"}`
		request := httptest.NewRequest(http.MethodGet, "/v1/crypto-upvote/users", strings.NewReader(stringfiedBody))
		writer := httptest.NewRecorder()

		repository := MockUsersRepository{}

		usersController := UsersController{
			UsersRepository: repository,
		}

		usersController.CreateUser(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)
		var body map[string]string
		json.Unmarshal(responseBody, &body)

		if !reflect.DeepEqual(http.StatusBadRequest, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusBadRequest, result.StatusCode)
		}
		if body["erro"] != "email: cannot be blank." {
			t.Errorf("wanted: %s, got: %s", "email: cannot be blank.", body["erro"])
		}
	})

	t.Run("CreateUser: Should return a error if dont pass name to register", func(t *testing.T) {
		stringfiedBody := `{"email":"email@email.com","password":"123456"}`
		request := httptest.NewRequest(http.MethodGet, "/v1/crypto-upvote/users", strings.NewReader(stringfiedBody))
		writer := httptest.NewRecorder()

		repository := MockUsersRepository{}

		usersController := UsersController{
			UsersRepository: repository,
		}

		usersController.CreateUser(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)
		var body map[string]string
		json.Unmarshal(responseBody, &body)

		if !reflect.DeepEqual(http.StatusBadRequest, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusBadRequest, result.StatusCode)
		}
		if body["erro"] != "name: cannot be blank." {
			t.Errorf("wanted: %s, got: %s", "name: cannot be blank.", body["erro"])
		}
	})

	t.Run("CreateUser: Should return a error if dont pass password to register", func(t *testing.T) {
		stringfiedBody := `{"name":"Unknow User","email":"email@email.com"}`
		request := httptest.NewRequest(http.MethodGet, "/v1/crypto-upvote/users", strings.NewReader(stringfiedBody))
		writer := httptest.NewRecorder()

		repository := MockUsersRepository{}

		usersController := UsersController{
			UsersRepository: repository,
		}

		usersController.CreateUser(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)
		var body map[string]string
		json.Unmarshal(responseBody, &body)

		if !reflect.DeepEqual(http.StatusBadRequest, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusBadRequest, result.StatusCode)
		}
		if body["erro"] != "password: cannot be blank." {
			t.Errorf("wanted: %s, got: %s", "password: cannot be blank.", body["erro"])
		}
	})

	t.Run("GetUser: Should return a registered user", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/v1/crypto-upvote/users/1", nil)
		writer := httptest.NewRecorder()

		vars := map[string]string{
			"userId": "1",
		}

		request = mux.SetURLVars(request, vars)

		repository := MockUsersRepository{}

		usersController := UsersController{
			UsersRepository: repository,
		}

		usersController.GetUser(writer, request)

		result := writer.Result()
		defer result.Body.Close()

		responseBody, _ := ioutil.ReadAll(result.Body)
		var body entities.User
		json.Unmarshal(responseBody, &body)

		log.Print(body)

		if !reflect.DeepEqual(http.StatusOK, result.StatusCode) {
			t.Errorf("wanted: %d, got: %d", http.StatusOK, result.StatusCode)
		}
		if body.Id != 1 {
			t.Errorf("wanted: %d, got: %d", 1, body.Id)
		}
		if body.Name != "Random User" {
			t.Errorf("wanted: %s, got: %s", "Random User", body.Name)
		}
		if body.Email != "email@email.com" {
			t.Errorf("wanted: %s, got: %s", "email@email.com", body.Email)
		}
		if body.Password != "" {
			t.Errorf("wanted: %s, got: %s", "", body.Password)
		}
	})
}
