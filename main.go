package main

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()
	r.GET("/healthcheck", handler)
	r.Run()
}

func handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "medically-core at your service!",
	})
}