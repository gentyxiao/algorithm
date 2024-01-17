package exercise

import "math"

// 寻找峰值-迭代方式实现，时间复杂度 O(n)
func FindPeakElement1(nums []int) int {
	idx := 0
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return idx
}

// 寻找峰值二分法实现，时间复杂度O(logn)
func FindPeakElement2(nums []int) int {
	n := len(nums)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, n-1
	for {
		mid := (left + right) / 2 // 等同 mid := left + (right - left)/2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
