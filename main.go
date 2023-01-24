package main

import (
	"fmt"
	"net/http"

	"go-link-shortener/handler"
	"go-link-shortener/store"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "This is a go link shortener",
			},
		)
	})

	router.POST(
		"/create-short-url", func(c *gin.Context) {
			handler.CreateShortUrl(c)

		},
	)

	router.GET(
		"/:shortUrl", func(c *gin.Context) {
			handler.HandleShortUrlRedirect(c)
		},
	)
	store.InitializeStore()
	err := router.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Errror: %v", err))
	}
}
