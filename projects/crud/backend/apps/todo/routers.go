package todo

import (
	"context"
	"fmt"
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest"
	"github.com/linksoft-dev/single/comps/go/requests"
	"net/http"
)

func getRoutes() *[]rest.Route {
	return &[]rest.Route{
		{Method: http.MethodGet, Path: "/todo/{id}", Handler: getTodo},
		{Method: http.MethodPost, Path: "/todo", Handler: createTodo},
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	id := r.URL.Query().Get("id")
	w.Write([]byte(fmt.Sprintf("Hello Go App framework, Id is '%s'", id)))
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	// parse the body into model
	m := Model{}
	if err := requests.ParseBody(r.Body, &m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create the service and save the object
	ctx := context.Background()
	service, err := NewService(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.Create(&m)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// return the response
	w.WriteHeader(http.StatusCreated)
	requests.ProcessResponse(w, m, ctx, err)
}
