package routes

import (
	"backend/src/controllers"
	"backend/src/repositories"
	"net/http"
)

var loginController controllers.LoginController = controllers.LoginController{
	UsersRepository: repositories.UserRepository{},
}

var loginRoutes = []Route{
	{
		URI:             "/v1/crypto-upvote/login",
		Method:          http.MethodPost,
		Function:        loginController.Login,
		IsAuthenticated: false,
	},
}
