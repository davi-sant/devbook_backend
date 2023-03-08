package router

import (
	"api/src/router/applicationRoutes"

	"github.com/gorilla/mux"
)

// (ToGenerate) -> vai retornar todas a rotas configuradas.
func ToGenerate() *mux.Router {
	router := mux.NewRouter()
	return applicationRoutes.ConfigRouters(router)
}
