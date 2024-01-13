package exercise

import "sort"

// [53]最大子数组和
func MaxSunArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	sort.Ints(dp)
	return dp[len(dp)-1]
}
