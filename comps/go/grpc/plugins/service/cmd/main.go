package main

import (
	genservice "github.com/linksoft-dev/single/comps/go/grpc/plugins/service"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		genservice.NewModule(),
	).RegisterPostProcessor(
		pgsgo.GoImports(),
		pgsgo.GoFmt(),
	).Render()
}
