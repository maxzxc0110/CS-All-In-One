# 描述
输入一个升序数组 array 和一个数字S，在数组中查找两个数，使得他们的和正好是S，如果有多对数字的和等于S，返回任意一组即可，如果无法找出这样的数字，返回一个空数组即可。

数据范围: 

 
0≤len(array)≤10^5


1≤array[i]≤10^6

  

# 暴力
```
func FindNumbersWithSum(array []int, sum int) []int {
	// write code here
	var ans []int
	for i, v := range array {
		var tmp_arr []int
		tmp_arr = array
		tmp_arr[i] = -100000  //把当前数字置为无效
		temp := sum - v  

		if IsContainInt(tmp_arr, temp) {  //判断两数之差是否在数组中
			ans = append(ans, v)
			ans = append(ans, temp)
			return ans
		}

	}
	return nil
}

func IsContainInt(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
```