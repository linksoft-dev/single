package todo

import (
	"github.com/graphql-go/graphql"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest"
	"google.golang.org/grpc"
	"net/http"
)

var App AppModule

type AppModule struct{}

func (a AppModule) GetRouters() *[]rest.Route {
	return getRoutes()
}

func (a AppModule) GetRouterGroup() *[]rest.RouteGroup {
	return nil
}

func (a AppModule) GetMiddlewares() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		middlewareLog,
	}
}

func (a AppModule) Register(s *grpc.Server, httpServer *runtime.ServeMux) error {
	return nil
}

func (a AppModule) BeforeStart() {}

func (a AppModule) AfterStart() {
}

func (a AppModule) GetInterceptor() grpc.UnaryServerInterceptor {
	return nil
}

func (a AppModule) GetGraphQLQueries() *graphql.Fields {
	return nil
}

func (a AppModule) GetGraphQLMutations() *graphql.Fields {
	return nil
}
