package utils

import (
	"log"
	"strconv"
	"strings"
)

//Calculator 计算器主体部分，通过辅助栈来快速计算结果
func Calculator(str string) int {
	stack := SimpleStack{}
	stack.Init()
	list := strings.Split(RemoveSpaces(str),"")
	NubFit(&list)
	num:=0
	for i, v := range list {
		if IsOperator(v) {
			if v == "-"{
				list[i] = "+"
				a, err := strconv.Atoi(list[i+1])
				if err!=nil{
					log.Fatal(err)
				}
				list[i+1] = strconv.Itoa(-a)
			}
			if v == "*" || v == "/"{
				num2, err := strconv.Atoi(list[i+1])
				if err!=nil{
				}
				list[i+1] = strconv.Itoa(Compute(num,list[i],num2))
				list[i] = "+"
				list[i-1] = "0"
			}
			num = 0
		}else {
			count, err := strconv.Atoi(list[i])
			if err!=nil{
				log.Fatal(err)
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
		if !IsOperator(v){
			sa, err := strconv.Atoi(v)
			if err!=nil{
				log.Fatal(err)
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
	return res
}









