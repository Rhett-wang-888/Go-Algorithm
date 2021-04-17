/**
 * @description: TODO 整数转换成罗马数字
 * @author rhett.wang
 * @date 2021/4/17 13:37
 * @version 1.0
 */
package main

func intToRoman(num int) string {
	dec := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	dir := make(map[int]string)
	dir[1000] = "M"
	dir[900] = "CM"
	dir[500] = "D"
	dir[400] = "CD"
	dir[100] = "C"
	dir[90] = "XC"
	dir[50] = "L"
	dir[40] = "XL"
	dir[10] = "X"
	dir[9] = "IX"
	dir[5] = "V"
	dir[4] = "IV"
	dir[1] = "I"
	ans := ""
	for _, v := range dec {
		if num/v != 0 {
			tmp := num / v
			for i := 0; i < tmp; i++ {
				ans += dir[v]
			}
			num = num - tmp*v
		}
	}
	return ans

}
