package main

import (
	"simple-calculator/calculatorTools"
	"testing"
)

func TestCalculator(t *testing.T)  {
	testStr := "2+3+7*1-7/2+12*2"
	res, err := calculatorTools.Calculator(testStr)
	if err != nil {
		t.Log(err.Error())
	}else{
		t.Log(res)
	}
}



