package model

import (
	"database/sql"
	"fmt"
	"math/rand"
)

func RecipeAll(db *sql.DB) ([]Recipe, error) {
	rows, err := db.Query(`select id,name,imageUrl from recipe`)
	if err != nil {
		return nil, err
	}
	return ScanRecipes(rows)
}

func RecipeRand(db *sql.DB) ([]Recipe, error) {
	rows, err := db.Query(`select id,name,imageUrl from recipe order by rand() limit 2`)
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
	query := fmt.Sprintf("select id,name,imageUrl from recipe order by %s desc limit 3", countColumn)

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return ScanRecipes(rows)
}

func RecipeLog(db *sql.DB, categoryId int, menuId string) error {

	// SQLインジェクション楽しい
	if !(1 <= categoryId && categoryId <= 10) {
		categoryId = rand.Intn(10) + 1
	}
	// 最高にひどい
	countColumn := fmt.Sprintf("count%d", categoryId)

	query := fmt.Sprintf("UPDATE recipe SET %s=%s+1 WHERE id=?", countColumn, countColumn)
	if _, err := db.Exec(query, menuId); err != nil {
		return err
	}
	return nil
}
