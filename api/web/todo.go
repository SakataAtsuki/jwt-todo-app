package web

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/SakataAtsuki/jwt-todo-app/model"
	"github.com/gin-gonic/gin"
)

func useTodos(e *gin.RouterGroup) {
	r := e.Group("/todos")

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
}
