/**
 * @description: TODO 选择排序
 * @author rhett.wang
 * @date 2022/4/22 10:24
 * @version 1.0
 */
package main

func main() {
	numbers := []int{6, 4, 3, 2, 7, 9, 8}
	SelectSort1(numbers)
	for i := 0; i < len(numbers); i++ {
		print(numbers[i])
	}

}
func SelectSort1(values []int) {
	length := len(values)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		min := i

		for j := length - 1; j > i; j-- {
			if values[j] < values[min] {
				min = j
			}
		}
		println("i:%d min:%d\n", i, min)
		values[i], values[min] = values[min], values[i]
		println(values)
	}
}

func SelectSort(list []int) {
	n := len(list)

	for i := 0; i < n-1; i++ {
		min := list[i]
		minIndex := i

		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}
