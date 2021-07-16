package router

import (
	"github.com/gin-gonic/gin"
	"simple-calculator/internal/ctrl"
	"simple-calculator/internal/globalError"
)

func Route(r *gin.Engine) *gin.Engine {
	//路由
	r.POST("/calculate",globalError.ErrorHandler(ctrl.Compute))
	return r
}
