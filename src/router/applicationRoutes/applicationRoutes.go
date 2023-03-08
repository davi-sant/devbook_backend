package applicationRoutes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// App router -> representa todas a rotas da api
type AppRouter struct {
	URI                    string
	Method                 string
	Func                   func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

// Config router -> adiciona todas a as rotas dentro do router
func ConfigRouters(r *mux.Router) *mux.Router {
	routers := userRoutes

	for _, router := range routers {
		r.HandleFunc(router.URI, router.Func).Methods(router.Method)
	}
	return r
}
