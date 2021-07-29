package utils

//Compute 两位运算
func Compute(a int, opt string, b int) int {
	switch opt {
	case "-":
		return sub(a, b)
	case "*":
		return mul(a, b)
	case "/":
		return div(a, b)
	default:
		return add(a, b)
	}
}

//真正的运算方法
func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func mul(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}
