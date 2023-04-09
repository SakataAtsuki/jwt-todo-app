package main

import (
	"github.com/SakataAtsuki/jwt-todo-app/model"
	"github.com/SakataAtsuki/jwt-todo-app/util"
	"github.com/SakataAtsuki/jwt-todo-app/web"
)

func init() {
	util.LoggingSettings("todoapp.log")
	model.DbConnect()
}

func main() {
	web.New()
}
