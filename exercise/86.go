package exercise

import "strings"

// 分割回文串
func Partition(s string) [][]string {
	// 1. 动态规划找出有哪些属于回文子串
	n := len(s)
	dp := make([][]bool, n)
	ans := [][]string{}
	sigle := []string{}
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		// 2. 单个字母都是回文子串
		sigle = append(sigle, string(s[i]))
	}
	ans = append(ans, sigle)
	item := []string{}
	// 3. 	其他回文子串（判断条件，相邻且相等属于回文子串，如果相等但是不相邻，但是内部串是
	// 是回文，则也是回文，即：acdca，a和a相等,cdc是回文则acdca也是回文）
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if i-j == 1 {
					dp[j][i] = true
				} else {
					dp[j][i] = dp[j+1][i-1]
				}
			}
			if dp[j][i] {
				// 存储回文子串
				item = append(item, s[j:i+1])
			}
		}
	}
	// 4. 字符串切分，按回文子串切分，找出切分处index，再将每个子串组合起来
	for _, v := range item {
		splits := []string{}
		idx := strings.Index(s, v)
		for i := 0; i < n; {
			if i >= idx && i < idx+len(v) {
				i += len(v)
				splits = append(splits, v)
				continue
			} else {
				splits = append(splits, string(s[i]))
				i++
			}

		}
		ans = append(ans, splits)
	}
	return ans
}
