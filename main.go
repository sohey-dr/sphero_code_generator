package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo"
	echomw "github.com/labstack/echo/middleware"
)

func main() {
    router := initRouter()
    router.Logger.Fatal(router.Start(":8080"))
}

func initRouter() *echo.Echo {
    e := echo.New()
    e.Use(echomw.CORSWithConfig(echomw.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowHeaders: []string{echo.HeaderContentType},
        AllowMethods: []string{echo.POST},
    }))
    e.POST("/sample", fileDownloadHandler)
    return e
}

func fileDownloadHandler(c echo.Context) error {
    file, err := os.Open("./test.js")
    defer file.Close()
    if err != nil {
        fmt.Println("error:file\n",err)
        return nil
    }

    setResponse(c, file)

    return c.NoContent(http.StatusOK)
}

func setResponse(c echo.Context, body io.Reader) {
    encodeName := url.QueryEscape("test.ts")
    response := c.Response()
    response.Header().Set("Cache-Control", "no-store")
    response.Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
    response.Header().Set(echo.HeaderAccessControlExposeHeaders, "Content-Disposition")
    response.Header().Set(echo.HeaderContentDisposition, "attachment; filename="+encodeName)
    response.WriteHeader(200)
    io.Copy(response.Writer, body)
    log.Println("response:", response)
}