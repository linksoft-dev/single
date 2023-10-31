package todo

import (
	"github.com/graphql-go/graphql"
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
	return nil
}

func (a AppModule) Register(_ *grpc.Server) {

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
