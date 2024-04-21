package exercise

// 移掉k位数字，使的新的字符串数字最小
func removeKdigits(num string, k int) string {
	// 思路：最靠左的位置数字最小
	// 相邻两个数左1左2，左1，与左2比较，如果左1大于左2，则删除左1
	// 解题：使用单调栈
	stack := []byte{}
	for i := range num {
		v := num[i]
		if stack[len(stack)-1] > v {
		}
		stack = append(stack, v)
	}
	return ""
}
