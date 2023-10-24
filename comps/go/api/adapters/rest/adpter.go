package rest

import (
	"github.com/graphql-go/graphql"
	"github.com/linksoft-dev/single/comps/go/api"
	"net/http"
)

// AppInterface interface that define the interface of App for Rest adapter
type AppInterface interface {
	api.App
	GetRouters() *[]Route
	GetRouterGroup() *[]RouteGroup
	GetMiddlewares() []func(http.Handler) http.Handler
	GetGraphQLQueries() *graphql.Fields
	GetGraphQLMutations() *graphql.Fields
}

// WebserverInterface interface that defines the adapter
type WebserverInterface interface {
	Run() error
	AddApp(app AppInterface)
	GetApps() []AppInterface
	GetName() string
}

type RouteGroup struct {
	Prefix  string
	Routers []Route
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}
