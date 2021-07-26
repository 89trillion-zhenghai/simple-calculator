package ctrl

import (
	"github.com/gin-gonic/gin"
	"simple-calculator/internal/globalError"
	"simple-calculator/internal/handler"
	"simple-calculator/internal/utils"
)

func Compute(c *gin.Context) (interface{},error) {

	exp := c.PostForm("exp")
	if len(exp) == 0 {
		//参数不为空校验
		return nil,globalError.ExpressionError("表达式为空",globalError.ExpIsEmpty)
	}
	if !utils.IsOtherChar(exp){
		return nil,globalError.ExpressionError("表达式含有非法字符",globalError.IllegalSymbol)
	}
	flag := utils.IsLegitimate(exp)
	if !flag{
		//表达式不合法
		return nil,globalError.ExpressionError("表达式不合法",globalError.InvalidExpression)
	}
	return handler.CalHandler(exp),nil
}
