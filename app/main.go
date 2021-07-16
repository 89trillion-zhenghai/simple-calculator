package main

import (
	"github.com/gin-gonic/gin"
	"simple-calculator/internal/router"
)

func main() {
	r := gin.Default()
	router.Route(r)
	r.Run(":8080")
}



