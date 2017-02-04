package api

import (
	"net/http"


	"github.com/labstack/echo"
)

func (r *Request) Test(c echo.Context) error {

	return c.JSON(http.StatusOK, "POST")
}
