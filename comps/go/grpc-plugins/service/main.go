package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func main() {

	// gen := pgs.Init(pgs.DebugEnv("DEBUG")).RegisterModule(module.New())

	// fmt.Println(gen.AST().Targets())
	// gen.Render()
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		NewModule(),
	).RegisterPostProcessor(
		pgsgo.GoFmt(),
	).Render()

}
