package main

import (
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
    e.POST("/", fileDownloadHandler)
    return e
}

func fileDownloadHandler(c echo.Context) error {
    fileName := "test.ts"
    file, err := fileOpen(fileName)
    if err != nil {
        return err
    }

    setResponse(c, file, fileName)

    return c.NoContent(http.StatusOK)
}

func fileOpen(fileName string) (*os.File, error) {
    file, err := os.Open(fileName)
    if err != nil {
        log.Println("error:file\n",err)
        return nil, err
    }
    defer file.Close()

    return file, nil
}

func setResponse(c echo.Context, body io.Reader, fileName string) {
    response := c.Response()
    response.Header().Set("Cache-Control", "no-store")
    response.Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
    response.Header().Set(echo.HeaderAccessControlExposeHeaders, "Content-Disposition")
    encodeName := url.QueryEscape(fileName)
    response.Header().Set(echo.HeaderContentDisposition, "attachment; filename="+encodeName)
    response.WriteHeader(200)
    io.Copy(response.Writer, body)
    log.Println("response:", response)
}