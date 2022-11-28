package router

import (
	"backend/src/router/routes"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigRouter(r)
}
