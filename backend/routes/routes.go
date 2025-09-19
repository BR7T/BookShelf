package routes

import (
	"github.com/gorilla/mux"
)
func InitRouter() *mux.Router{
	router := mux.NewRouter()
	setupUserRoutes(router)
	setupBookRoutes(router)
	return  router
}
