package routes

import (
	"backend/src/controllers"
	"backend/src/repositories"
	"net/http"
)

var cryptosController controllers.CryptoController = controllers.CryptoController{
	CryptoRepository: repositories.CryptoRepository{},
}

var cryptosRoutes = []Route{
	{
		URI:             "/v1/crypto-upvote/cryptos",
		Method:          http.MethodGet,
		Function:        cryptosController.GetCryptos,
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{cryptoId}",
		Method:          http.MethodGet,
		Function:        cryptosController.GetCrypto,
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{cryptoId}/up",
		Method:          http.MethodPut,
		Function:        cryptosController.Vote,
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{cryptoId}/down",
		Method:          http.MethodPut,
		Function:        cryptosController.Vote,
		IsAuthenticated: false,
	},
}
