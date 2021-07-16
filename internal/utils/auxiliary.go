package utils

import (
	"strconv"
	"strings"
)

//Compute 两位运算
func Compute(a int,opt string,b int) int {
	switch opt {
	case "-":
		return sub(a,b)
	case "*":
		return mul(a,b)
	case "/":
		return div(a,b)
	default:
		return add(a,b)
	}
}
//真正的运算方法
func add(a int,b int) int {
	return a + b
}
func sub(a int,b int) int {
	return a - b
}
func mul(a int,b int) int {
	return a * b
}
func div(a int,b int) int {
	return a / b
}

//RemoveSpaces 移除空格
func RemoveSpaces(str string) string{
	cpStr := make([]string,len(str))
	for _, v := range strings.Split(str,"") {
		if v != " "{
			cpStr = append(cpStr, v)
		}
	}
	return strings.Join(cpStr,"")
}

//NubFit 将连续的数值字符串连接 eg  "2","3","+","3"  - >   "23","+","3"
func NubFit(str *[]string) {
	num := 0
	copy := make([]string,0)
	for _,v:= range *str{
		count, err := strconv.Atoi(v)
		if err != nil {
			copy = append(copy, strconv.Itoa(num))
			copy = append(copy, v)
			num = 0
		}else{
			num = num * 10 +count
		}
	}
	copy = append(copy, strconv.Itoa(num))
	*str = copy
}