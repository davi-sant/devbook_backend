package router

import "github.com/gorilla/mux"

// (ToGenerate) -> vai retornar todas a rotas configuradas.
func ToGenerate() *mux.Router {
	return mux.NewRouter()
}
