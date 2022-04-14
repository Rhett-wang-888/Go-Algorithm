/**
 * @description: TODO
 * @author rhett.wang
 * @date 2022/4/14 17:37
 * @version 1.0
 */
package main

func main() {
	a := []int{20, 10, 30, 50, 40, 60}
	b := shellsortArray(a)

	println(len(b))
	for i := 0; i < len(b); i++ {
		print(b[i], " ")
	}
}
func shellsortArray(nums []int) []int {
	gap := len(nums) / 2

	for gap > 0 {
		for i := gap; i < len(nums); i++ {
			j := i
			for j-gap >= 0 && nums[j-gap] > nums[j] {
				nums[j-gap], nums[j] = nums[j], nums[j-gap]
				j -= gap
			}
		}
		gap /= 2
	}
	return nums
}
