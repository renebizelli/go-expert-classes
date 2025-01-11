package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO categories (id, name, description) values (?, ?, ?)", id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {

	rows, err := c.db.Query("SELECT id, name, description FROM categories")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []Category{}

	for rows.Next() {

		category := Category{}

		if e := rows.Scan(&category.ID, &category.Name, &category.Description); e != nil {
			return nil, e
		}

		categories = append(categories, category)

	}

	return categories, nil
}

func (c *Category) FindByCourseId(courseId string) (*Category, error) {

	rows := c.db.QueryRow("SELECT a.id, a.name, a.description FROM categories a INNER JOIN courses o ON o.categoryId = a.id WHERE o.id = ?", courseId)

	category := Category{}

	if e := rows.Scan(&category.ID, &category.Name, &category.Description); e != nil {
		return nil, e
	}

	return &category, nil
}
