package server

import (
	"os"
	"todoapp/src/middleware"
	"todoapp/src/http"
	"todoapp/src/factory"

	"github.com/labstack/echo/v4"
)

func Init() {
	var PORT = os.Getenv("PORT")
	e := echo.New()
	middleware.Init(e)
	f := factory.NewFactory()
	http.Init(e, f)
	e.Logger.Fatal(e.Start(":" + PORT))
}