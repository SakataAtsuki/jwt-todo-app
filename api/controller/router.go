package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/SakataAtsuki/jwt-todo-app/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"DELETE",
			"PUT",
		},
		AllowHeaders: []string{
			"Authorization",
		},
	}))

	v1 := r.Group("todo/api/v1")
	v1.Use(authMiddleware())
	{
		v1.GET("/todos", todosGET)
		v1.POST("/todos", todoPOST)
		v1.PATCH("/todos/:id", todoPATCH)
		v1.DELETE("/todos/:id", todoDELETE)
	}
	r.Run(":8080")
}

func todosGET(c *gin.Context) {
	todos, err := model.GetTodos()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func todoPOST(c *gin.Context) {
	title := c.PostForm("title")
	err := model.CreateTodo(title)
	if err != nil {
		log.Fatalln(err)
	}
}

func todoPATCH(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	todo, err := model.GetTodo(uint(id))
	if err != nil {
		log.Fatalln(err)
	}

	title := c.PostForm("title")
	now := time.Now()
	todo.Title = title
	todo.UpdatedAt = now
	todo.UpdateTodo()
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func todoDELETE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatalln(err)
	}
	todo, err := model.GetTodo(uint(id))
	if err != nil {
		log.Fatalln(err)
	}
	err = todo.DeleteTodo()
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, "Deleted")
}
