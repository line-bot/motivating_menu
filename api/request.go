package api

import (
	"net/http"

	"database/sql"

	"github.com/labstack/echo"
)

type Request struct {
	DB *sql.DB
}

func (r *Request) GetTest(c echo.Context) error {
	return c.JSON(http.StatusOK, "GET")
}

func (r *Request) PostTest(c echo.Context) error {
	return c.JSON(http.StatusOK, "POST")
}

func (r *Request) PutTest(c echo.Context) error {
	return c.JSON(http.StatusOK, "PUT")
}

func (r *Request) DeleteTest(c echo.Context) error {
	return c.JSON(http.StatusOK, "DELETE")
}
