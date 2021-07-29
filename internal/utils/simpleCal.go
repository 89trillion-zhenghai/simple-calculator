package utils

import (
	"strconv"
	"strings"
)

//Calculator 计算器主体部分，通过辅助栈来快速计算结果
func Calculator(str string) int {
	stack1 := SimpleStack{}
	stack1.Init()
	chs := make([]string, 0)
	exp := strings.Split(str, "")
	exp = append(exp, "+", "0")
	num := 0
	for _, v := range exp {
		i, err := strconv.Atoi(v)
		if err == nil {
			num = num*10 + i
		} else if IsOperator(v) {
			stack1.Push(num)
			num = 0
			if len(chs) > 0 {
				last := chs[len(chs)-1]
				if last == "*" || last == "/" {
					pop1 := stack1.Pop()
					pop2 := stack1.Pop()
					stack1.Push(Compute(pop2, last, pop1))
					chs = chs[:len(chs)-1]
				}
			}
			chs = append(chs, v)
		}
	}
	stack1.Push(num)
	stack2 := SimpleStack{}
	for !stack1.IsEmpty() {
		stack2.Push(stack1.Pop())
	}
	for _, ch := range chs {
		pop1 := stack2.Pop()
		pop2 := stack2.Pop()
		stack2.Push(Compute(pop1, ch, pop2))
	}
	return stack2.Pop()
}
