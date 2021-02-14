/**
 * @description: TODO
 * @author rhett.wang
 * @date 2021/2/14 14:38
 * @version 1.0
 */
package main

import "fmt"

func main() {
	fmt.Println(LengthOfLongestSubstring("eeeseseses"))
}
func LengthOfLongestSubstring(s string) int {
	flag := [256]int{}
	beg := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if flag[s[i]] > 0 && flag[s[i]] > beg {
			beg = flag[s[i]]
		}
		flag[s[i]] = i + 1
		max = maxnum(max, i-beg+1)
	}
	return max

}
func maxnum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
