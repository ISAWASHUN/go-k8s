package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
  log.Println("Start Server")
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	log.Fatal(r.Run())
}
