package utils

//SimpleStack 简易版栈，利用切片实现了栈的基础功能入栈和出栈
type SimpleStack struct {
	nodes []int
}

//Init 初始化栈
func (stack *SimpleStack) Init()  {
	stack.nodes = make([]int,0)
}

//IsEmpty 判断栈是否为空
func (stack SimpleStack) IsEmpty() bool {
	return len(stack.nodes) == 0
}

//Push 入栈
func (stack *SimpleStack) Push(node int) {
	stack.nodes = append(stack.nodes,node)
}

//Pop 出栈 返回栈内最顶部的元素
func (stack *SimpleStack) Pop() (node int) {
	cpStack := *stack
	top := cpStack.nodes[len(cpStack.nodes)-1]
	stack.nodes = cpStack.nodes[:len(cpStack.nodes)-1]
	return top
}
func (stack SimpleStack) Len() int {
	return len(stack.nodes)
}

