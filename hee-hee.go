package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Quotes struct {
	quote  string
	source string
}

var quotes = []string{
	"All of us are products of our childhood.",
	"To give someone a piece of your heart, is worth more than all the wealth in the world.",
	"A star can never die. It just turns into a smile and melts back into the cosmic music, the dance of life.",
	"I'm just like anyone. I cut and I bleed and I embarass easily.",
	"Look beyond yourself...",
	"Success, fame, and fortune, they're all illusions. All there is that is real is the friendship that two can share.",
}

func AllMJQuotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quotes)
}


func main() {
	router := gin.Default()

	router.GET("/", AllMJQuotes)
	router.Run("localhost:8000")
}
