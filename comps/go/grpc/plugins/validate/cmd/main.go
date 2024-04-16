package main

import (
	genservice "github.com/linksoft-dev/single/comps/go/grpc/plugins/validate"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		genservice.NewModule(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()
}
