package phaseTwo

import (
	"database/sql"
	"net/http"

	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/line-bot/motivating_menu/model"
)

type PhaseTwo struct {
	DB *sql.DB
}

func (p *PhaseTwo) Response(c echo.Context) error {
	recipe, err := model.RecipeAll(p.DB)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, recipe)
}
