package exercise

// 55. 跳跃游戏
func canJump(nums []int) bool {
	// 任意一个位置i，有i+nums[i] 大于等于j,那么该位置j可以到达
	// 以此类推，局部求解实现最终求解
	rightMost, n := 0, len(nums)
	for i := 0; i < n; i++ {
		if i <= rightMost {
			// 表示当前可到达的最右位置
			// i是当前遍历的位置
			rightMost = max(rightMost, nums[i]+i)
			if rightMost >= n-1 {
				return true
			}
		}
	}
	return false
}
