package rest

import (
	"github.com/waykiss/wkgo"
	"net/http"
)

// AppInterface interface that define the interface of App for Rest adapter
type AppInterface interface {
	wkgo.App
	GetRouters() *[]Route
	GetRouterGroup() *[]RouteGroup
	GetMiddlewares() []func(http.Handler) http.Handler
}

// WebserverInterface interface that defines the adapter
type WebserverInterface interface {
	Run()
	Add(app AppInterface)
	GetApps() []wkgo.App
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
