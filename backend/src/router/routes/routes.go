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

func ConfigRouter(router *mux.Router) *mux.Router {
	routes := cryptosRoutes
	routes = append(routes, imagesRoutes...)
	routes = append(routes, usersRoutes...)
	routes = append(routes, loginRoutes...)

	for _, route := range routes {
		if route.IsAuthenticated {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return router
}
