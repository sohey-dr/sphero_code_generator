package handler

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo"
)

func FileDownloadHandler(c echo.Context) error {
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