package model

import (
	"go.course.todo-api/views"
)

func Read(name string) ([]views.PostRequest, error) {
	if len(name) > 0 {
		return ReadByName(name)
	}
	return ReadAll()
}

func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM todo WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil
}

func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO todo VALUES (?, ?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodo(name string) error {
	deleteQ, err := con.Query("DELETE FROM todo WHERE name=?", name)
	defer deleteQ.Close()
	if err != nil {
		return err
	}
	return nil
}
