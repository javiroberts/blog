package errors

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NotFound(c echo.Context) error {
	e := ErrorSchema{
		Code:    http.StatusNotFound,
		Reason:  "resource_not_found",
		Message: "The requested resource was not found in the server.",
	}
	fmt.Println(e)
	return c.Render(e.Code, "error", e)
}
