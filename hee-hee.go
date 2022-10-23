package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var quotes = []string{
	"HEE HEE",
	"SHAMONE",
	"AOOW",
}

func AllMJQuotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quotes)
}


func main() {
	router := gin.Default()

	router.GET("/", AllMJQuotes)
	router.Run("localhost:6969")
}
