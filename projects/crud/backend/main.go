package main

import (
	"github.com/linksoft-dev/single/comps/go/api"
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest/chi"
	"github.com/linksoft-dev/single/crud/apps/todo"
)

func main() {

	chiServer := chi.New("8085", "/chi")
	chiServer.AddApp(todo.App)

	api.AddAdapter(chiServer)
	api.Start("crud-sample")
}
