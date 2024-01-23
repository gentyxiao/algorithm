package exercise

// 回文子串的数量
func CountSubstrings(s string) int {
	n := len(s)
	ans := n
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i-j == 1 {
					dp[j][i] = true
				} else {
					dp[j][i] = dp[j+1][i-1]
				}
				if dp[j][i] {
					ans++
				}
			}
		}
	}
	return ans
}
