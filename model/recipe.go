package model

import "database/sql"

func RecipeAll(db *sql.DB) ([]Recipe, error) {
	rows, err := db.Query(`select * from recipe`)
	if err != nil {
		return nil, err
	}
	return ScanRecipes(rows)
}
