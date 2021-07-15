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
	list := strings.Split(RemoveSpaces(str),"")
	nubFit(&list)
	num:=0
	for i, v := range list {
		if isOperator(v) {
			if v == "-"{
				list[i] = "+"
				a, err := strconv.Atoi(list[i+1])
				if err!=nil{
				}
				list[i+1] = strconv.Itoa(-a)
			}
			if v == "*" || v == "/"{
				num2, err := strconv.Atoi(list[i+1])
				if err!=nil{
				}
				list[i+1] = strconv.Itoa(compute(num,list[i],num2))
				list[i] = "+"
				list[i-1] = "0"

			}
			num = 0
		}else {
			count, err := strconv.Atoi(list[i])
			if err!=nil{

			}
			if num != 0 {
				list[i-1] = "0"
			}
			if num < 0 {
				num = num * 10 - count
			}else{
				num = num * 10 + count
			}
			list[i] = strconv.Itoa(num)
		}
	}

	for _,v := range list {
		if !isOperator(v){
			sa, err := strconv.Atoi(v)
			if err != nil {

			}
			stack.Push(sa)
		}
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

	//判断v是否是运算符 是返回true
	func isOperator(v string) bool {
		if v == "+" || v == "-" || v == "*" || v == "/"{
			return true
		}
		return false
	}

	//IsLegitimate 判断运算表达式是否合法，合法返回true，不合法返回false
	func IsLegitimate(str string) bool {
		s := strings.Split(str, "")
		if isOperator(s[0]) || isOperator(s[len(s)-1]){
			return false
		}
		flag := 0
		for _,v:=range s{
			if isOperator(v){
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
func nubFit(str *[]string) {
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