package chi

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/linksoft-dev/single/comps/go/api"
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var webserver *chi.Mux

func New(port, prefix string) *Adapter {
	return &Adapter{port: port, prefix: prefix}
}

// Adapter struct to save apps that implement grpc adapters and set port that grpc server should run
type Adapter struct {
	port   string
	prefix string
	apps   []rest.AppInterface
}

func (g *Adapter) AddApp(app rest.AppInterface) {
	g.apps = append(g.apps, app)
}

func (g *Adapter) Run() error {
	if webserver == nil {
		webserver = chi.NewRouter()

		for _, app := range g.apps {
			appMiddleware := app.GetMiddlewares()
			if appMiddleware != nil {
				for _, value := range appMiddleware {
					if value != nil {
						webserver.Use(value)
					}
				}
			}

			restRouters := app.GetRouters()
			if restRouters != nil {
				for _, route := range *restRouters {
					if g.prefix != "" {
						route.Path = fmt.Sprintf("%s/%s", g.prefix, strings.TrimPrefix(route.Path, "/"))
					}
					webserver.MethodFunc(route.Method, route.Path, route.Handler)
					log.Infof("%s - Adding route %v", g.GetName(), map[string]interface{}{"Route": "/api" + route.Path})
				}
			}

			queries := app.GetGraphQLQueries()
			if queries != nil {
				for key, value := range *queries {
					rootQuery.AddFieldConfig(key, value)
				}
			}

			mutations := app.GetGraphQLMutations()
			if mutations != nil {
				for key, value := range *mutations {
					rootMutation.AddFieldConfig(key, value)
				}
			}
		}
		go func() {
			http.ListenAndServe(":"+g.port, webserver)
		}()

		// test server connection
		for {
			_, err := http.Get("http://localhost:" + g.port)
			if err == nil {
				log.Infof("Chi server listening on %s\n", g.port)
				return nil
			}
		}
	}
	return nil
}

func (g *Adapter) GetName() string {
	return "Chi WebServer"
}

func (g Adapter) GetApps() []api.App {
	apps := []api.App{}
	for _, app := range g.apps {
		apps = append(apps, app)
	}
	return apps
}
