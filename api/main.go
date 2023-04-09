package main

import (
	"github.com/SakataAtsuki/jwt-todo-app/controller"
	"github.com/SakataAtsuki/jwt-todo-app/model"
)

func init() {
	// utils.LoggingSettings("todoapp.log")
	model.DbConnect()
}

func main() {
	controller.StartServer()
}
