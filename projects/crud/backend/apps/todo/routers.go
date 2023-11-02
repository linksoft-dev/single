package todo

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest"
	"net/http"
)

func getRoutes() *[]rest.Route {
	return &[]rest.Route{
		{Method: http.MethodGet, Path: "/todo/{id}", Handler: getTodo},
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("Hello Go App framework, Id is '%s'", id)))
}
