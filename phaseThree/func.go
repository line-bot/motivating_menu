package phaseThree

import (
	"database/sql"
	"net/http"

	"strconv"

	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/line-bot/motivating_menu/model"
)

type PhaseThree struct {
	DB *sql.DB
}

func (p *PhaseThree) Response(c echo.Context) error {
	categoryId := c.QueryParam("category_id")
	menuId := c.QueryParam("menu_id")
	ci, _ := strconv.Atoi(categoryId)
	err := model.RecipeLog(p.DB, ci, menuId)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, "ok")
}
