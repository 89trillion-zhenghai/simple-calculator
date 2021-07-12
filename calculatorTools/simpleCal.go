package calculatorTools

import (
	"errors"
	"simple-calculator/model"
	"strconv"
	"strings"
)

//Calculator 计算器主体部分，通过辅助栈来快速计算结果
func Calculator(str string) (result int ,err error) {
	if !IsLegitimate(str){
		return -1,errors.New("不合法的算术表达式，请检查！")
	}
	stack := model.SimpleStack{}
	stack.Init()
	s := "+"
	num := 0
	for i,v := range strings.Split(str, ""){
		c, err := strconv.Atoi(v)
		if err == nil {
			num = num * 10 + c
		}else if v != " " || i == len(str){
			if s == "+"{
				stack.Push(num)
			} else if s == "-" {
				stack.Push(-num)
			}else {
				stack.Push(compute(stack.Pop(),s,num))
			}
			num = 0
			s = v
		}
	}
	if s == "+"{
		stack.Push(num)
	} else if s == "-" {
		stack.Push(-num)
	}else{
		stack.Push(compute(stack.Pop(),s,num))
	}

	res := 0
	for {
		if stack.IsEmpty(){
			break
		}
		res += stack.Pop()
	}
	return res,nil
}


//Compute 两位运算
func compute(a int,opt string,b int) int {
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

//判断切片str是否存在运算符 存在返回false
func isOperator(str ... string) bool {
	for _,v:=range str{
		if v == "+" || v == "-" || v == "*" || v == "/"{
			return false
		}
	}
	return true
}

//IsLegitimate 判断运算表达式是否合法，合法返回true，不合法返回false
func IsLegitimate(str string) bool {
	s := strings.Split(str, "")
	if !isOperator(s[0],s[len(s)-1]){
		return false
	}
	flag := 0
	for _,v:=range s{
		if !isOperator(v){
			flag ++
			if flag == 2{
				return false
			}
		}else if v != " "{
			flag  = 0
		}
	}
	return true
}
