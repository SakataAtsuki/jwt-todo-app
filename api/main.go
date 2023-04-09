package main

import (
	"github.com/SakataAtsuki/jwt-todo-app/controller"
	"github.com/SakataAtsuki/jwt-todo-app/model"
	"github.com/SakataAtsuki/jwt-todo-app/util"
)

func init() {
	util.LoggingSettings("todoapp.log")
	model.DbConnect()
}

func main() {
	controller.StartServer()
}
