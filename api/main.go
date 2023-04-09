package main

import (
	"github.com/SakataAtsuki/jwt-todo-app/api/controller"
	"github.com/SakataAtsuki/jwt-todo-app/api/model"
)

func init() {
	// utils.LoggingSettings("todoapp.log")
	model.DbConnect()
}

func main() {
	controller.StartServer()
}
