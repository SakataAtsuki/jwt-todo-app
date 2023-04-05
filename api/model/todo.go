package model

import (
	"log"
	"time"
)

type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetTodo(id uint) (todo Todo, err error) {
	err = Db.QueryRow("SELECT id, title, created_at, updated_at FROM todos WHERE id = ?", id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)
	if err != nil {
		log.Fatalln(err)
	}
	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	rows, err := Db.Query("SELECT id, title, created_at, updated_at FROM todos")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.CreatedAt,
			&todo.UpdatedAt,	
		)
		todos = append(todos, todo)
	}
	return todos, err
}

func CreateTodo(title string) (err error) {
	now := time.Now()
	_, err = Db.Exec("INSERT INTO todos (title, created_at, updated_at) VALUES(?, ?, ?)", title, now, now)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (todo *Todo) UpdateTodo() error {
	_, err = Db.Exec("UPDATE todos SET title = ? WHERE id = ?", todo.Title, todo.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (todo *Todo) DeleteTodo() error {
	_, err = Db.Exec("DELETE FROM todos WHERE id = ?", todo.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
