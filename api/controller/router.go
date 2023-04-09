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
	e := gin.Default()

	e.Use(cors.New(cors.Config{
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

	r := e.Group("todo/api/v1/todos")

	r.Use(authMiddleware())

	r.GET("", func(c *gin.Context) {
		todos, err := model.GetTodos()
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{"todos": todos})
	})

	r.POST("", func(c *gin.Context) {
		title := c.PostForm("title")
		err := model.CreateTodo(title)
		if err != nil {
			log.Fatalln(err)
		}
	})

	r.PATCH(":id", func(c *gin.Context) {
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
	})

	r.DELETE(":id", func(c *gin.Context) {
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
	})

	e.Run(":8080")
}
