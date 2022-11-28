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
		Function:        func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Up vote")) },
		IsAuthenticated: true,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{cryptoId}/down",
		Method:          http.MethodPut,
		Function:        func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Down vote")) },
		IsAuthenticated: true,
	},
}
