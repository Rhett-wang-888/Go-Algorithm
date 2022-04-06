/**
 * @description: TODO
 * @author rhett.wang
 * @date 2022/4/6 15:04
 * @version 1.0
 */
package main

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

func main() {
	a := []int{20, 10, 30, 50, 40, 60}
	b := sortArray(a)
	print(len(b))
	for i := 0; i < len(b); i++ {
		println(b[i])
	}

}
