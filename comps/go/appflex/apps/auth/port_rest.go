package auth

import (
	"github.com/waykiss/wkgo/adapters/rest"
	"net/http"
)

//getRoutes return the rest routers related
func getRoutes() *[]rest.RouteGroup {
	return &[]rest.RouteGroup{
		{
			Prefix: "users", Routers: []rest.Route{
				{Method: http.MethodGet, Path: "/", Handler: list},
				{Method: http.MethodPost, Path: "/register", Handler: register},
				{Method: http.MethodPost, Path: "/login", Handler: login},
				{Method: http.MethodDelete, Path: "/", Handler: deleteUser},
			},
		},
	}
}

func getMiddlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		validateTokenMiddleware,
	}
}

func validateTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPublicRouter(r.Method, r.RequestURI) {
			next.ServeHTTP(w, r)
			return
		}
		token := r.Header.Get("token")
		err := validateToken(token)
		if err != nil {
			rest.Response(w, nil, err)
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

//register handle to receive register request
func register(w http.ResponseWriter, r *http.Request) {
	m := Model{}
	if err := rest.GetRequestParams(r, &m); err != nil {
		rest.Response(w, nil, err)
		return
	}
	service, err := NewService()
	if err != nil {
		rest.Response(w, nil, err)
		return
	}
	valErr := service.register(m.Name, m.Email, m.Password, m.Age)
	rest.Response(w, nil, valErr)
}

func list(w http.ResponseWriter, r *http.Request) {
	service, err := NewService()
	if err != nil {
		rest.Response(w, nil, err)
		return
	}
	users, valErr := service.Find(Query{})
	rest.Response(w, users, valErr)
}

// login api rest referente a login
func login(w http.ResponseWriter, r *http.Request) {
	m := loginModel{}
	if err := rest.GetRequestParams(r, &m); err != nil {
		rest.Response(w, nil, err)
		return
	}
	service, err := NewService()
	if err != nil {
		rest.Response(w, nil, err)
		return
	}
	result, valErr := service.login(m.Email, m.Password)
	rest.Response(w, result, valErr)
}

// login api rest referente a login
func deleteUser(w http.ResponseWriter, r *http.Request) {
	m := Model{}
	if err := rest.GetRequestParams(r, &m); err != nil {
		rest.Response(w, nil, err)
		return
	}
	service, err := NewService()
	if err != nil {
		rest.Response(w, nil, err)
		return
	}
	valErr := service.Delete(m.Id)
	rest.Response(w, "ok", valErr)
}
