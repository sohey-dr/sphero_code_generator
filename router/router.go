package router

import (
  "sphero_code_generator/handler"

	"github.com/labstack/echo"
	echomw "github.com/labstack/echo/middleware"
)

func InitRouter() *echo.Echo {
    e := echo.New()
    e.Use(echomw.CORSWithConfig(echomw.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{echo.HeaderContentType},
        AllowMethods: []string{echo.POST},
    }))
    e.POST("/", handler.FileDownloadHandler)
    return e
}
