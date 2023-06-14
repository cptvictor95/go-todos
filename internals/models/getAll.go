package models

import (
	"fmt"

	"github.com/cptvictor95/go-todos/internals/database/db"
)

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {
			fmt.Printf("error getting row: %v", err)
			continue
		}

		todos = append(todos, todo)
	}

	return
}
