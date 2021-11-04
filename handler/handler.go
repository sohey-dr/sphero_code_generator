package handler

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"sphero_code_generator/service"

	"github.com/labstack/echo"
)

type FileDownloadRequest struct {
    Programs []string `json:"programs"`
}

func FileDownloadHandler(c echo.Context) error {
    var programs FileDownloadRequest
    if err := c.Bind(&programs); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    err := service.GenerateCode()
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
    file, err := os.Open("sphero.ts")
    if err != nil {
        return err
    }
    io.Copy(response.Writer, file)

    return c.NoContent(http.StatusOK)
}
