package fiber

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/graphql-go/handler"
	"github.com/linksoft-dev/single/comps/go/api"
	"github.com/linksoft-dev/single/comps/go/api/adapters/rest"
	log "github.com/sirupsen/logrus"
)

var fiberApp *fiber.App

func New(port, prefix string, config *fiber.Config) *Adapter {
	if config == nil {
		config = &fiber.Config{
			BodyLimit: 40 * 1024 * 1024, // 40 mb  response limit as default
		}
	}
	return &Adapter{port: port, prefix: prefix, config: *config}
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

var routes map[string]http.HandlerFunc

func (g *Adapter) Run() error {
	routes = map[string]http.HandlerFunc{}
	if fiberApp == nil {
		fiberApp = fiber.New(g.config)

		apiGroup := fiberApp.Group(g.prefix)
		staticGroup := fiberApp.Group("/")

		fiberAddInternalMiddlewares(apiGroup)
		g.fiberSpaMiddleware(staticGroup)

		for _, app := range g.apps {
			appMiddleware := app.GetMiddlewares()
			for _, value := range appMiddleware {
				if value != nil {
					apiGroup.Use(adaptor.HTTPMiddleware(value))
				}
			}

			restRouters := app.GetRouters()
			if restRouters != nil {
				for _, route := range *restRouters {
					route.Path = convertBraceToColon(route.Path)
					apiGroup.Add(route.Method, route.Path, adaptor.HTTPHandlerFunc(route.Handler))
					//routes[route.Path] = route.Handler
					//
					//apiGroup.Add(route.Method, route.Path, func(c *fiber.Ctx) error {
					//	r := getRequestFromFiberContext(c)
					//	res := NewCustomResponseWriter()
					//	path := strings.ReplaceAll(c.Route().Path, g.prefix, "")
					//	handlerFunc := routes[path]
					//	if handlerFunc != nil {
					//		handlerFunc(res, r)
					//	}
					//	c.Write(res.body)
					//	return c.SendString(string(res.body))
					//})

					log.Infof("%s - Adding route %v", g.GetName(), map[string]interface{}{"Route": g.prefix + route.Path})
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

func convertBraceToColon(path string) string {
	return strings.ReplaceAll(strings.ReplaceAll(path, "{", ":"), "}", "")
}

func (g Adapter) GetApps() []api.App {
	apps := []api.App{}
	for _, app := range g.apps {
		apps = append(apps, app)
	}
	return apps
}

// getRequestFromFiberContext this function return a pointer of http.Request copying URL parameters
// given fiber context
func getRequestFromFiberContext(c *fiber.Ctx) (r *http.Request) {
	r, _ = adaptor.ConvertRequest(c, true)
	params := c.AllParams()
	if len(params) == 0 {
		return
	}
	// copy URL parameters
	queryParams := url.Values{}
	for key, value := range params {
		queryParams.Add(key, value)
	}
	r.URL.RawQuery = queryParams.Encode()
	return
}

type CustomResponseWriter struct {
	body       []byte
	statusCode int
	header     http.Header
}

func NewCustomResponseWriter() *CustomResponseWriter {
	return &CustomResponseWriter{
		header: http.Header{},
	}
}

func (w *CustomResponseWriter) Header() http.Header {
	return w.header
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body = b
	// implement it as per your requirement
	return 0, nil
}

func (w *CustomResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
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

func (g Adapter) fiberSpaMiddleware(group fiber.Router) {

	group.Use(filesystem.New(filesystem.Config{
		Root:         http.Dir("./web"),
		Browse:       true,
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
		Next: func(c *fiber.Ctx) bool {
			if strings.Contains(c.Request().URI().String(), g.prefix) {
				return true
			}
			return false
		},
	}))
}
