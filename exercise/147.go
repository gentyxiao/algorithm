package exercise

import "math"

// 最小栈
type MinStackWrap struct {
	Stack    []int
	MinStack []int
}

func ContructorMinStack() *MinStackWrap {
	return &MinStackWrap{
		Stack:    make([]int, 0),
		MinStack: []int{math.MaxInt32},
	}
}

func (m *MinStackWrap) Push(x int) {
	m.Stack = append(m.Stack, x)
	m.MinStack = append(m.MinStack, min(m.MinStack[len(m.MinStack)-1], x))
}

func (m *MinStackWrap) Pop() {
	m.Stack = m.Stack[:len(m.MinStack)-1]
	m.MinStack = m.MinStack[:len(m.MinStack)-1]
}

func (m *MinStackWrap) Top() int {
	return m.Stack[len(m.MinStack)-1]
}

func (m *MinStackWrap) GetMin() int {
	return m.MinStack[len(m.MinStack)-1]
}
