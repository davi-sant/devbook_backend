package applicationRoutes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []AppRouter{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Func:                   controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Func:                   controllers.GetUsers,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Func:                   controllers.GetUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Func:                   controllers.UpdateUser,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Func:                   controllers.DeleteUser,
		RequiresAuthentication: false,
	},
}
