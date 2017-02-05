package phaseTwo

import (
	"database/sql"
	"net/http"

	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/line-bot/motivating_menu/model"
	"strconv"
)

type PhaseTwo struct {
	DB *sql.DB
}

func (p *PhaseTwo) Response(c echo.Context) error {
	categoryId := c.QueryParam("category_id")
	ci, _ := strconv.Atoi(categoryId)

	recipe1, err := model.RecipeRecommend(p.DB, ci)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	recipe2, err := model.RecipeRand(p.DB)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	var recipe []model.Recipe
	recipe = append(recipe1, recipe2...)

	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, recipe)
}
