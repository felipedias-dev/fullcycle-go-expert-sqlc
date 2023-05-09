package main

import (
	"context"
	"database/sql"

	"github.com/felipedias-dev/fullcycle-go-expert-sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "TI",
		Description: sql.NullString{String: "TI Courses", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		ID:          "0ed4ac5d-daba-4a95-9486-7b060c1a731c",
		Name:        "TI Updated",
		Description: sql.NullString{String: "TI Courses Updated", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, cacategory := range categories {
		println(cacategory.ID, cacategory.Name, cacategory.Description.String)
	}

	err = queries.DeleteCategory(ctx, "0ed4ac5d-daba-4a95-9486-7b060c1a731c")
	if err != nil {
		panic(err)
	}
}
