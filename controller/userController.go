package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-calculator/calculatorTools"
)

func GetCalculator(c *gin.Context)  {
	msg := ""
	str := c.PostForm("str")
	fmt.Println(str)
	result, err := calculatorTools.Calculator(str)
	if err != nil {
		msg = err.Error()
	}else{
		msg = "success!"
	}
	c.JSON(http.StatusOK,gin.H{
		"res" : result,
		"message" : msg,
	})
}
