package applicationRoutes

import "net/http"

var userRoutes = []AppRouter{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Func:                   func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Func:                   func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Func:                   func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuthentication: false,
	},
	{
		URI:                    "users/{userId}",
		Method:                 http.MethodDelete,
		Func:                   func(w http.ResponseWriter, r *http.Request) {},
		RequiresAuthentication: false,
	},
}
