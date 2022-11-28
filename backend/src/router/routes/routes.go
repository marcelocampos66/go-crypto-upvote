package routes

import (
	"backend/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI             string
	Method          string
	Function        func(http.ResponseWriter, *http.Request)
	IsAuthenticated bool
}

func ConfigRouter(r *mux.Router) *mux.Router {
	routes := cryptosRoutes

	for _, route := range routes {
		r.HandleFunc(
			route.URI,
			middlewares.Logger(route.Function),
		).Methods(route.Method)
	}

	return r
}
