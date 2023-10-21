package main

import (
	"github.com/kissprojects/single/comps/go/api"
	"github.com/kissprojects/single/comps/go/api/adapters/rest/fiber"
	"github.com/kissprojects/single/comps/go/api/apps/auth"
)

func main() {
	restAdapter := fiber.New("8000")
	api.AddAdapters(restAdapter)

	api.AddApp(auth.App)
	api.Start("crud-sample")
}
