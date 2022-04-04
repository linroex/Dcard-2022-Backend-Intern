package main

import (
	"github.com/gin-gonic/gin"
)

type ShortUrlRequest struct {
	Url      string
	ExpireAt string
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/urls", urls)
	}

	router.GET("/:id", goUrl)

	router.Run("127.0.0.1:8080")

}
