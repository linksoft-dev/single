package main

import (
	"github.com/linksoft-dev/single/comps/go/api"
	"github.com/linksoft-dev/single/comps/go/api/adapters/grpc"
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest"
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest/chi"
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest/fiber"
	"github.com/linksoft-dev/single/crud/apps/todo"
)

func main() {

	type allApps interface {
		grpc.AppI
		rest.AppInterface
	}
	apps := []allApps{
		todo.App,
	}

	chiServer := chi.New("8085", "/chi")
	fiberServer := fiber.New("8086", "/api", nil)

	for _, app := range apps {
		chiServer.AddApp(app)
		fiberServer.AddApp(app)
	}
	api.AddAdapter(chiServer)
	api.AddAdapter(fiberServer)
	api.Start("crud-sample")
}
