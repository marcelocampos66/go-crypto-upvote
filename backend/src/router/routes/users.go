package routes

import (
	"backend/src/controllers"
	"backend/src/repositories"
	"net/http"
)

var usersController controllers.UsersController = controllers.UsersController{
	UsersRepository: repositories.UserRepository{},
}

var usersRoutes = []Route{
	{
		URI:             "/v1/crypto-upvote/users",
		Method:          http.MethodPost,
		Function:        usersController.CreateUser,
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/users/{userId}",
		Method:          http.MethodGet,
		Function:        usersController.GetUser,
		IsAuthenticated: true,
	},
}
