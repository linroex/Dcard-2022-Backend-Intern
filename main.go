package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type ShortUrlRequest struct {
	Url      string `binding:"required"`
	ExpireAt string
}

type Env struct {
	db *sql.DB
}

func main() {
	db, _ := sql.Open("mysql", os.Getenv("dbConnectString"))
	defer db.Close()
	app := &Env{
		db: db,
	}

	router := setupRouter(app)
	router.Run("127.0.0.1:8080")
}

func setupRouter(app *Env) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/urls", app.urls)
	}

	router.GET("/:id", app.goUrl)

	return router
}
