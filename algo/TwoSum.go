/**
 * @description: TODO 简单的计算2个数之和，返回一个数组
 * @author rhett.wang
 * @date 2021/8/30 11:38
 * @version 1.0
 */
package main

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	hash := map[int]int{}
	for i, x := range nums {
		if p, ok := hash[target-x]; ok {
			return []int{p, i}
		}
		hash[x] = i
	}
	return nil
}
