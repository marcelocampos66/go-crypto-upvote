package routes

import (
	"net/http"
)

var cryptosRoutes = []Route{
	{
		URI:             "/v1/crypto-upvote/cryptos",
		Method:          http.MethodGet,
		Function:        func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Get Cryptos!")) },
		IsAuthenticated: false,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{userId}/up",
		Method:          http.MethodPut,
		Function:        func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Up vote")) },
		IsAuthenticated: true,
	},
	{
		URI:             "/v1/crypto-upvote/cryptos/{userId}/down",
		Method:          http.MethodPut,
		Function:        func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Down vote")) },
		IsAuthenticated: true,
	},
}
