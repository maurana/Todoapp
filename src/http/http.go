package http

import (
	"os"
	"fmt"
	"net/http"
	"todoapp/src/factory"
	"todoapp/src/app/list"
	"todoapp/src/app/sublist"
	docs "todoapp/docs"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

// @title Todo App API
// @version 0.0.1
// @description Todo App Server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /
func Init(e *echo.Echo, f *factory.Factory) {
	var (
		APP     = os.Getenv("APP")
		VERSION = os.Getenv("VERSION")
		HOST    = os.Getenv("BASE_URL")
		SCHEME  = os.Getenv("SCHEME")
	)

	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to REST API of %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	docs.SwaggerInfo.Title = APP
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = HOST
	docs.SwaggerInfo.Schemes = []string{SCHEME}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes
	list.NewHandler(f).Route(e.Group("/list"))
	sublist.NewHandler(f).Route(e.Group("/sublist"))
}