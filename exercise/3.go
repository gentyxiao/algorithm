package exercise

// [3]无重复字符的最长子串的长度
// 滑动窗口
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	right := 0
	hashMap := map[byte]int{}
	ans := 0
	for left := 0; left < n; left++ {
		if left != 0 {
			delete(hashMap, s[left-1])
		}
		for ; right < n && hashMap[s[right]] == 0; right++ {
			hashMap[s[right]]++
			ans = max(ans, right-left+1)
		}
	}
	return ans
}
