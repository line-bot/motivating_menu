package api

import (
	"net/http"

	"database/sql"

	"github.com/labstack/echo"
)

type Request struct {
	DB *sql.DB
}

func (r *Request) Test(c echo.Context) error {

	return c.JSON(http.StatusOK, "POST")
}
