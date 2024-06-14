package category

import (
	"fmt"
	"os"

	"github.com/dystrow/budget/internal"
	"github.com/labstack/gommon/log"
)

const tablename string = "category"

func InitDB() {
	conn, err := internal.DBConnect()
	if err != nil {
		log.Fatalf("initdb: failed while trying to connect to db: %w", err)
		os.Exit(1)
	}

	create := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL, goal REAL)", tablename)
	if _, err := conn.Exec(create); err != nil {
		log.Fatalf("initdb: failed to execute query: %w", err)
		os.Exit(1)
	}
}

func dbDelete(id string) error {
	conn, err := internal.DBConnect()
	if err != nil {
		return fmt.Errorf("delete category: failed while trying to connect to db: %w", err)
	}
	defer conn.Close()

	query := fmt.Sprintf("DELETE FROM %s WHERE id = '%s'", tablename, id)
	_, err = conn.Exec(query)
	if err != nil {
		return fmt.Errorf("delete category: failed while executing query: %w", err)
	}

	return nil
}

func dbAdd(name string) (int64, error) {
	conn, err := internal.DBConnect()
	if err != nil {
		return -1, fmt.Errorf("add category: failed while trying to connect to db: %w", err)
	}
	defer conn.Close()

	query := fmt.Sprintf("INSERT INTO %s(name) VALUES ('%s')", tablename, name)
	res, err := conn.Exec(query)
	if err != nil {
		return -1, fmt.Errorf("add category: failed while executing query: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("add category: failed to extract id. caution: category might have been created %w", err)
	}

	return id, nil
}

func dbGetAll() ([]Category, error) {
	conn, err := internal.DBConnect()
	if err != nil {
		return []Category{}, fmt.Errorf("get all categories: failed while trying to connect to db: %w", err)
	}
	defer conn.Close()

	query := fmt.Sprintf("SELECT * FROM %s", tablename)
	rows, err := conn.Query(query)
	if err != nil {
		return []Category{}, fmt.Errorf("get all categories: failed while executing query: %w", err)
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var cat Category

		rows.Scan(&cat.Id, &cat.Name, &cat.Goal)
		categories = append(categories, cat)
	}

	if err := rows.Err(); err != nil {
		return []Category{}, fmt.Errorf("get all categories: error while iterationg through response: %w", err)
	}

	return categories, nil
}

func dbGetById(id string) (Category, error) {
	conn, err := internal.DBConnect()
	if err != nil {
		return Category{}, fmt.Errorf("get category by id: failed while trying to connect to db: %w", err)
	}
	defer conn.Close()

	query := fmt.Sprintf("SELECT * FROM %s where id='%s'", tablename, id)
	rows, err := conn.Query(query)
	if err != nil {
		return Category{}, fmt.Errorf("get category by id: failed while executing query: %w", err)
	}
	defer rows.Close()

	categories := []Category{}
	for rows.Next() {
		var cat Category

		rows.Scan(&cat.Id, &cat.Name, &cat.Goal)
		categories = append(categories, cat)
	}

	if err := rows.Err(); err != nil {
		return Category{}, fmt.Errorf("get category by id: Error while iterationg through response: %w", err)
	}
	if len(categories) > 1 {
		return Category{}, fmt.Errorf("get category by id: Returned multiple Categories for given ID: %w", err)
	}

	return categories[0], nil
}
