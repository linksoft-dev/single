package fiber

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/waykiss/wkgo"
	"github.com/waykiss/wkgo/adapters/rest"
	"log"
	"net/http"
	"strings"
)

//adapter private structure that implement rest adapter
type adapter struct {
	port string
	apps []rest.AppInterface
}

var webserver *fiber.App

const apiPrefix = "/api"

func New(port string) rest.WebserverInterface {
	return &adapter{port: port}
}

//Run method to start fiber webserver
func (f adapter) Run() {
	if webserver == nil {
		// configs disponívels
		webserver = fiber.New(
			fiber.Config{
				//TODO this should be paramatized
				BodyLimit: 40 * 1024 * 1024, // 40 mb for upload limit size
			})

		apiGroup := webserver.Group(apiPrefix)
		for _, app := range f.apps {
			appmid := app.GetMiddlewares()
			if appmid != nil {
				for _, value := range appmid {
					if value != nil {
						apiGroup.Use(adaptor.HTTPMiddleware(value))
					}
				}
			}

			restRouters := app.GetRouters()
			if restRouters != nil {
				for _, route := range *restRouters {
					apiGroup.Add(route.Method, route.Path, adaptor.HTTPHandlerFunc(route.Handler))
					log.Printf("Adding route %s", route.Path)
				}
			}

			restRouterGroups := app.GetRouterGroup()
			if restRouterGroups != nil {
				for _, routeGroup := range *restRouterGroups {
					fiberGroup := apiGroup.Group(routeGroup.Prefix)
					log.Printf("Adding routes for group :%s", routeGroup.Prefix)
					for _, route := range routeGroup.Routers {
						fiberGroup.Add(route.Method, route.Path, adaptor.HTTPHandlerFunc(route.Handler))
						log.Printf("Adding route : %s %s%s", route.Method, routeGroup.Prefix, route.Path)
					}
				}
			}

		}
		log.Fatal(webserver.Listen(":" + f.port))
	}
}

func (f adapter) GetApps() (r []wkgo.App) {
	for _, app := range f.apps {
		r = append(r, app)
	}
	return
}

func (f *adapter) Add(app rest.AppInterface) {
	f.apps = append(f.apps, app)
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
