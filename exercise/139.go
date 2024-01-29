package exercise

// 单词拆分:动态规划
func WordBreak(s string, wordDict []string) bool {
	hashMap := map[string]bool{}
	for _, v := range wordDict {
		hashMap[v] = true
	}
	dp := make([]bool, len(s)+1)
	// 空单词视为可以被拆分组成
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// 如果单词j之前组成都可以拆分为wordDict中的单词
			// 且单词j-i部分如果也能被可以拆分为wordDict中的单词
			// 那么i之前的所有单词都可以可以拆分为wordDict中的单词
			if dp[j] && hashMap[string(s[j:i])] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
