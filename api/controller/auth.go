package controller

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		opt := option.WithCredentialsFile(os.Getenv("GOOGLE_CREDENTIALS_JSON"))

		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Printf("err: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Printf("err: %v\n", err)
			os.Exit(1)
		}

		authHandler := c.Request.Header.Get("Authorization")
		idToken := strings.Replace(authHandler, "Bearer ", "", 1)

		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Error": err.Error(),
			})
			return
		}
		log.Printf("Vertifed ID token: %v\n", token)
		c.Next()
	}
}
