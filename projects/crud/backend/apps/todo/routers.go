package todo

import (
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest"
	"net/http"
)

func getRoutes() *[]rest.Route {
	return &[]rest.Route{
		{Method: http.MethodGet, Path: "/todo", Handler: getTodo},
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Chi"))
}
