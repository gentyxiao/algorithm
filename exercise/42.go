package exercise

// 接雨水
// 动态规划
func Trap(height []int) int {
	n := len(height)
	rightDp, leftDp := make([]int, n), make([]int, n)
	leftDp[0] = height[0]
	rightDp[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		leftDp[i] = max(leftDp[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		rightDp[i] = max(rightDp[i+1], height[i])
	}
	ans := 0
	for i, h := range height {
		ans += min(leftDp[i], rightDp[i]) - h
	}
	return ans
}
