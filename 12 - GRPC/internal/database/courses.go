package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryId string) (Course, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO courses (id, name, description, categoryId) values (?, ?, ?, ?)", id, name, description, categoryId)

	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description, CategoryId: categoryId}, nil
}

func (c *Course) FindAll() ([]Course, error) {

	rows, err := c.db.Query("SELECT id, name, description, categoryId FROM courses")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}

	for rows.Next() {

		course := Course{}

		if e := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId); e != nil {
			return nil, e
		}

		courses = append(courses, course)

	}

	return courses, nil
}

func (c *Course) FindByCategoryId(categoryId string) ([]Course, error) {

	rows, err := c.db.Query("SELECT id, name, description, categoryId FROM courses WHERE categoryId = ?", categoryId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}

	for rows.Next() {

		course := Course{}

		if e := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryId); e != nil {
			return nil, e
		}

		courses = append(courses, course)

	}

	return courses, nil
}
