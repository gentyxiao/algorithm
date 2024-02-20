package exercise

import "strconv"

// 字符串相乘，模拟乘法算法过程
func multiply(num1 string, num2 string) string {
	len1, len2 := len(num1), len(num2)
	if len1 == 0 || len2 == 0 || num1 == "0" || num2 == "0" {
		return "0"
	}
	// 两数相乘的积的长度最大不会超过两数长度之和
	result := make([]int, len1+len2)
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			x := int(num1[i]) - '0'
			y := int(num2[j]) - '0'
			r := x * y
			sum := result[i+j+1] + r // 加上当前位乘积之和
			result[i+j+1] = sum % 10 // 当前位取模
			result[i+j] += sum / 10  // 进位相加
		}
	}
	ans := ""
	for k, v := range result {
		if k == 0 && v == 0 {
			continue
		}
		ans += strconv.Itoa(v)
	}
	return ans
}
