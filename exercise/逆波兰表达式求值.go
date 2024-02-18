package exercise

import (
	"strconv"
)

// 逆波兰表达式求值
func EvalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		n, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, n)
		} else {
			n1, n2, n3 := stack[len(stack)-2], stack[len(stack)-1], 0
			switch token {
			case "+":
				n3 = n1 + n2
			case "-":
				n3 = n1 - n2
			case "*":
				n3 = n1 * n2
			case "/":
				n3 = n1 / n2
			}
			stack = stack[:len(stack)-2] // 切片范围拷贝是左开右闭
			stack = append(stack, n3)
		}
	}
	return stack[0]
}
