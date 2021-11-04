package handler

import (
	"io"
	"net/http"
	"net/url"
    "sphero_code_generator/service"

	"github.com/labstack/echo"
)

func FileDownloadHandler(c echo.Context) error {
    file, err := service.GenerateCode()
    if err != nil {
        return err
    }

    response := c.Response()
    response.Header().Set("Cache-Control", "no-store")
    response.Header().Set(echo.HeaderContentType, echo.MIMEOctetStream)
    response.Header().Set(echo.HeaderAccessControlExposeHeaders, "Content-Disposition")
    encodeName := url.QueryEscape("sphero.ts")
    response.Header().Set(echo.HeaderContentDisposition, "attachment; filename="+encodeName)
    response.WriteHeader(200)
    io.Copy(response.Writer, file)

    return c.NoContent(http.StatusOK)
}
