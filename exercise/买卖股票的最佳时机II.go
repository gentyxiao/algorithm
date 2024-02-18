package exercise

// 买卖股票的最佳时机 II
func maxProfit2(prices []int) int {
	// 动态规划
	// dp[i][0/1] 表示不持有股票或持有股票获得的最大利益
	// dp[i][0] = max{前一天不持有股票，前一天持有股票+今日卖出股票}
	// dp[i][1] = max{前一天持有股票，前一天不持有股票-今日买入股票}
	// dp[0][0] = 0
	// dp[0][1] = 0-prices[0]
	dp := make([][2]int, len(prices))
	dp[0][1] = -prices[0]
	dp[0][0] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}

func maxProfit2x(prices []int) int {
	// 贪心算法
	// 由于不限制买卖次数，只需满足不能买入前必须卖出即可
	// 所以只需要考虑前日股票价格和今日股票价格
	// 当今日股票大于昨日股票价格即可以卖出获得收益
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += max(0, prices[i]-prices[i-1])
	}
	return ans
}
