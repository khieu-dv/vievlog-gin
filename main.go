package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Dinh nghia router
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// khoi chay server tren cong 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Fail to start server: %v", err)
	}
}
