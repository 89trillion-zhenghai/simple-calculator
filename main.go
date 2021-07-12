package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simple-calculator/controller"
)

func main() {
	r := gin.Default()
	r.POST("calculator", controller.GetCalculator)
	r.Run(":8080")
	fmt.Println("xxxx")
	fmt.Println("xxxx")
}




