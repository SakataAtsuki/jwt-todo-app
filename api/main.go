package main

import (
	"jwt-todo-app/api/controller"
	"jwt-todo-app/api/model"
)

func init() {
	// utils.LoggingSettings("todoapp.log")
	model.DbConnect()
}

func main() {
	controller.StartServer()
}
