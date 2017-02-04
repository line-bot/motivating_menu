package phaseTwo

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

type PhaseTwo struct {
	DB *sql.DB
}

var res = `
[
    {
        "category": "1",
        "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/hamburger_240.jpeg"
    },
    {
        "category": "2",
        "imageUrl": "https://s3-us-west-2.amazonaws.com/lineapitest/pizza_240.jpeg"
    },
    {
        "category": "3",
        "https://s3-us-west-2.amazonaws.com/lineapitest/bread_240.jpeg"
    }
]
`

func (p *PhaseTwo) Response(c echo.Context) error {
	return c.JSON(http.StatusOK, res)
}
