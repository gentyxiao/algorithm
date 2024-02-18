package exercise

import "math"

// 121. 买卖股票的最佳时机
func maxProfit(prices []int) int {
	// 贪心算法，局部最优，实现最终最优
	cost := math.MaxInt32
	got := 0
	for _, v := range prices {
		got = max(got, v-cost)
		cost = min(cost, v)
	}
	return got
}
