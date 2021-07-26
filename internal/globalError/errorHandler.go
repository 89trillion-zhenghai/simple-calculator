package globalError

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	InvalidExpression  = 1001	//表达式不合法
	IllegalSymbol	   = 10011	//出现非法字符
	HeadOperator	   = 10012	//表达式首字符为运算符
	EndOperator		   = 10013	//表达式最后一个字符为运算符
	SuccessiveOperator = 10014  //连续出现运算符
	ExpIsEmpty		   = 10015	//表达式为空

)

type GlobalHandler func(c *gin.Context) (interface{}, error)


func ErrorHandler(handler GlobalHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := handler(c)
		if err != nil {
			globalError := err.(GlobalError)
			c.JSON(globalError.Status, globalError)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"计算结果": data,
		})
	}
}

// ExpressionError 表达式异常
func ExpressionError(message string, code int) GlobalError {
	return GlobalError{
		Status:  http.StatusBadRequest,
		Code:    code,
		Message: message,
	}
}



