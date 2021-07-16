package handler

import "simple-calculator/internal/utils"

func CalHandler(exp string) int {
	//调用计算工具进行计算
	result := utils.Calculator(exp)
	return result
}
