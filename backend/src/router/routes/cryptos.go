package routes

import (
	"backend/src/controllers"
	"backend/src/repositories"
	"net/http"
)

var controller controllers.CryptoController = controllers.CryptoController{
	CryptoRepository: repositories.CryptoRepository{},
}

var cryptosRoutes = []Route{
	{
		URI:             "/v1/crypto-upvote/cryptos",
		Method:          http.MethodGet,
		Function:        controller.GetCryptos,
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{cryptoId}",
		Method:          http.MethodGet,
		Function:        controller.GetCrypto,
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{cryptoId}/up",
		Method:          http.MethodPut,
		Function:        controller.Vote,
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{cryptoId}/down",
		Method:          http.MethodPut,
		Function:        controller.Vote,
		IsAuthenticated: false,
	},
}
