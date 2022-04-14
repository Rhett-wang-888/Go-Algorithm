/**
 * @description: TODO
 * @author rhett.wang
 * @date 2022/4/6 15:04
 * @version 1.0
 */
package main

import (
	"fmt"
	"math/rand"
)

func insertSort4(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		key := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j -= 1
		}
		nums[i+1] = key
	}
	fmt.Println(nums)
	return nums
}
func insertSort3(nums []int) []int {
	length := len(nums)
	for j := 1; j < length; j++ {
		key := nums[j]
		i := j - 1

		for i >= 0 && nums[i] > key {
			nums[i+1] = nums[i]
			i -= 1
		}
		nums[i+1] = key

	}
	fmt.Println(nums)
	return nums
}
func main() {
	var length = 10
	var tree []int
	for i := 0; i < length; i++ {
		tree = append(tree, int(rand.Intn(100)))
	}
	fmt.Println(tree)
	for i := 1; i < length; i++ {
		for j := i; j > 0 && tree[j] < tree[j-1]; j-- {
			tree[j], tree[j-1] = tree[j-1], tree[j]
		}
		fmt.Println(tree)
	}

	//a := []int{20, 10, 30, 50, 40, 60}
	//b := sortArray(a)
	//print(len(b))
	//for i := 0; i < len(b); i++ {
	//	println(b[i])
	//}

}

func sortArray(nums []int) []int {
	// 插入排序，比较交换，稳定算法，时间O(n^2)，空间O(1)
	// 0->len方向，每轮从后往前比较，相当于找到合适位置，插入进去
	// 数据规模小的时候，或基本有序，效率高
	for i := 0; i < len(nums); i++ {
		// 从后往前比较，找到位置插入
		for j := i; j > 0; j-- {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}
func sortArray1(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		tmp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > tmp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = tmp
	}
	return nums
}
