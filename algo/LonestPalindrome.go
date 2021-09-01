/**
 * @description: TODO 最长回文字符串
 * @author rhett.wang
 * @date 2021/9/1 15:52
 * @version 1.0
 */
package main

func centerExpend(m int, n int, s string, ans string) (bool, string) {
	for m > 0 && n < len(s) && s[m] == s[n] {
		m--
		n++
	}
	if len(ans) < n-m-1 {
		return true, s[m+1 : n]
	}
	return false, ans
}
func longestPalindrome(s string) string {

	var ret string = ""
	var need_save bool = false
	var save_str string = ""

	if len(s) < 2 {
		return s
	}
	for i := 0; i < len(s); i++ {
		if need_save, save_str = centerExpend(i, i, s, ret); need_save == true {
			ret = save_str
		}
		if need_save, save_str = centerExpend(i, i+1, s, ret); need_save == true {
			ret = save_str
		}

	}
	return ret
}

func longestPalindrome1(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	maxString := ""
	maxLength := 0
	var maxSubString func(left, right int)
	maxSubString = func(left, right int) {
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		length := right - left - 1
		if length > maxLength {
			maxLength = length
			maxString = s[left+1 : right]
		}
	}

	for i := 0; i < n; i++ {
		maxSubString(i, i)
		maxSubString(i, i+1)
	}

	return maxString
}
func longestPalindrome2(s string) string {
	var ans string
	var n = len(s)
	var dp = make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ {
		for i := 0; i+1 < n; i++ {
			j := i + 1
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[j] == s[i] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] == 1 && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}
