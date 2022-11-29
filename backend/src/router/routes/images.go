package routes

import (
	"backend/src/controllers"
	"net/http"
)

var imagesController controllers.ImagesController = controllers.ImagesController{}

var imagesRoutes = []Route{
	{
		URI:             "/v1/crypto-upvote/images/{cryptoCode}",
		Method:          http.MethodGet,
		Function:        imagesController.GetImage,
		IsAuthenticated: false,
	},
}
