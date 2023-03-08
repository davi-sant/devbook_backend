package applicationRoutes

import "net/http"

// App router -> representa todas a rotas da api
type AppRouter struct {
	URI                    string
	Method                 string
	Func                   func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}
