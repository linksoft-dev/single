package main

import (
	"github.com/kissprojects/single/comps/go/appflex"
	"github.com/kissprojects/single/comps/go/appflex/adapters/rest/fiber"
)

func main() {
	//grpcAdapter := grpc.New("1559")
	//grpcAdapter.Add(auth.App)
	restAdapter := fiber.New("8000")
	appflex.AddAdapters(restAdapter)
	appflex.AddApp(auth.App)
	appflex.Start("crud-sample")
}
