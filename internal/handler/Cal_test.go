package handler

import (
	"simple-calculator/internal/utils"
	"testing"
)

func TestSimpleStack(t *testing.T) {

	stack1 := utils.SimpleStack{}
	stack1.Push(1)
	stack1.Push(2)
	println(stack1.Pop())
	println(stack1.Pop())
}

func TestCalculator(t *testing.T) {

	exp := "7 +  9/3 * 24  + 10"
	result := utils.Calculator(exp)
	println(result)
}
