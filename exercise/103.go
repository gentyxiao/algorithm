package exercise

import "math"

// 零钱兑换，完全背包问题
func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		// 第一层，遍历背包
		for j := 0; j < len(coins); j++ {
			// 第二层，遍历物品
			if coins[j] <= i {
				// 当前背包能放下该物品，则数量+1，同时加上剩余容量
				// 放不下就取dp[i]的值
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
