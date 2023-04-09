package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New() {
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

	r := e.Group("todo/api/v1")

	useTodos(r)

	e.Run(":8080")
}
