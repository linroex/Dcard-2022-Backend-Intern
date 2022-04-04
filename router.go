package main

import (
	"database/sql"
	"net/url"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

func urls(c *gin.Context) {
	db, _ := sql.Open("mysql", os.Getenv("dbConnectString"))

	request := ShortUrlRequest{}

	c.BindJSON(&request)

	_, err := url.ParseRequestURI(request.Url)
	if err != nil {
		c.JSON(300, gin.H{
			"message": "url invalid",
		})
	}

	stmt, _ := db.Prepare("INSERT INTO records set url=?,slug=? ;")

	slug := RandString(3)
	res, _ := stmt.Exec(request.Url, slug)
	lastId, _ := res.LastInsertId()

	c.JSON(200, gin.H{
		"id":       strconv.FormatInt(lastId, 10) + slug,
		"shortUrl": os.Getenv("baseUrl") + strconv.FormatInt(lastId, 10) + slug,
	})

	db.Close()
}
