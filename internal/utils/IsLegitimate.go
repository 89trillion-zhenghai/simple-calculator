package utils

import "strings"

//IsLegitimate 判断运算表达式是否合法，合法返回true，不合法返回false
func IsLegitimate(str string) bool {
	s := strings.Split(str, "")
	if IsOperator(s[0]) || IsOperator(s[len(s)-1]){
		return false
	}
	flag := 0
	for _,v:=range s{
		if IsOperator(v){
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

//IsOperator 判断参数v是否是运算符 是返回true
func IsOperator(v string) bool {
	if v == "+" || v == "-" || v == "*" || v == "/"{
		return true
	}
	return false
}