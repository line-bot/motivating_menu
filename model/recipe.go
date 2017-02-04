package model

import (
	"database/sql"
	"fmt"
	"math/rand"
)

func RecipeAll(db *sql.DB) ([]Recipe, error) {
	rows, err := db.Query(`select id,name from recipe`)
	if err != nil {
		return nil, err
	}
	return ScanRecipes(rows)
}

func RecipeRecommend(db *sql.DB, categoryId int) ([]Recipe, error) {

	// SQLインジェクション楽しい
	if !(1 <= categoryId && categoryId <= 10) {
		categoryId = rand.Intn(10) + 1
	}

	// 最高にひどい
	countColumn := fmt.Sprintf("count%d", categoryId)
	//countColumn := "count" + categoryId
	query := fmt.Sprintf("select id,name from recipe order by %s desc limit 5", countColumn)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return ScanRecipes(rows)
}
