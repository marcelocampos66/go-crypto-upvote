package routes

import (
	"backend/src/controllers"
	"net/http"
)

var cryptosRoutes = []Route{
	{
		URI:             "/v1/crypto-upvote/cryptos",
		Method:          http.MethodGet,
		Function:        controllers.GetCryptos,
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{cryptoId}/up",
		Method:          http.MethodPut,
		Function:        controllers.Vote,
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{cryptoId}/down",
		Method:          http.MethodPut,
		Function:        controllers.Vote,
		IsAuthenticated: false,
	},
}
