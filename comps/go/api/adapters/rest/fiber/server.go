package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/graphql-go/handler"
	"github.com/kissprojects/single/comps/go/api"
	"github.com/kissprojects/single/comps/go/api/adapters/rest"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var fiberApp *fiber.App

func New(port, prefix string, config fiber.Config) *Adapter {
	return &Adapter{port: port, prefix: prefix, config: config}
}

// Adapter struct to save apps that implement grpc adapters and set port that grpc server should run
type Adapter struct {
	port   string
	prefix string
	apps   []rest.AppInterface
	config fiber.Config
}

func (g *Adapter) AddApp(app rest.AppInterface) {
	g.apps = append(g.apps, app)
}

const apiPrefix = "/api"

func (g *Adapter) Run() error {

	if fiberApp == nil {
		fiberApp = fiber.New(g.config)

		apiGroup := fiberApp.Group(apiPrefix)
		staticGroup := fiberApp.Group("/")

		fiberAddInternalMiddlewares(apiGroup)
		fiberSpaMiddleware(staticGroup)

		for _, app := range g.apps {
			appMiddleware := app.GetMiddlewares()
			if appMiddleware != nil {
				for _, value := range appMiddleware {
					if value != nil {
						apiGroup.Use(adaptor.HTTPMiddleware(value))
					}
				}
			}

			restRouters := app.GetRouters()
			if restRouters != nil {
				for _, route := range *restRouters {
					apiGroup.Add(route.Method, route.Path, adaptor.HTTPHandlerFunc(route.Handler))
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
		createGraphQlSchema()
		fiberGraphQLHandler(apiGroup)
		go func() {
			log.Fatal(fiberApp.Listen(":" + g.port))
		}()

		// test server connection
		for {
			_, err := http.Get("http://localhost:" + g.port)
			if err == nil {
				log.Infof("Fiber server listening on %s\n", g.port)
				return nil
			}
		}
	}
	return nil
}

func (g *Adapter) GetName() string {
	return "Fiber WebServer"
}

func (g Adapter) GetApps() []api.App {
	apps := []api.App{}
	for _, app := range g.apps {
		apps = append(apps, app)
	}
	return apps
}

// fiberGraphQLHandler cria o handler para processar requests graphQl
func fiberGraphQLHandler(apiGroup fiber.Router) {

	// Se houver um schema de graphql definido, crie o handler http com o schema
	if schema != nil {
		h := handler.New(&handler.Config{
			Schema:     schema,
			Pretty:     true,
			Playground: true,
		})

		apiGroup.Get("/graphql", adaptor.HTTPHandler(h))
		apiGroup.Post("/graphql", adaptor.HTTPHandler(h))
	}
}

func fiberAddInternalMiddlewares(apiGroup fiber.Router) {
	// CORS
	apiGroup.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "org,token,access-control-allow-origin, access-control-allow-headers, content-type," +
			"access-control-allow-methods",
		AllowCredentials: true,
	}))

	// Logs
	apiGroup.Use(logger.New(logger.Config{
		Format:   "${pid} ${status} - ${method} ${path}\n",
		TimeZone: "America/New_York",
	}))
	apiGroup.Use(recover.New(recover.ConfigDefault))

	// middleware para pegar o host da API
	//apiGroup.Use(func(c *fiber.Ctx) (err error) {
	//	if AppConfig.Host == "" {
	//		AppConfig.Host = fmt.Sprintf("%s://%s", c.Protocol(), c.Hostname())
	//	}
	//	if AppConfig.ApiHost == "" {
	//		AppConfig.ApiHost = fmt.Sprintf("%s%s", AppConfig.Host, apiPrefix)
	//	}
	//
	//	if AppConfig.FrontendHost == "" {
	//		AppConfig.FrontendHost = fmt.Sprintf("%s", AppConfig.Host)
	//	}
	//
	//	return c.Next()
	//})

}

func fiberSpaMiddleware(group fiber.Router) {

	group.Use(filesystem.New(filesystem.Config{
		Root:         http.Dir("./web"),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
		Next: func(c *fiber.Ctx) bool {
			if strings.Contains(c.Request().URI().String(), apiPrefix) {
				return true
			}
			return false
		},
	}))
}
