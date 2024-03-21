# 描述
小明很喜欢数学,有一天他在做数学作业时,要求计算出9~16的和,他马上就写出了正确答案是100。但是他并不满足于此,他在想究竟有多少种连续的正数序列的和为100(至少包括两个数)。没多久,他就得到另一组连续正数和为100的序列:18,19,20,21,22。现在把问题交给你,你能不能也很快的找出所有和为S的连续正数序列?

数据范围：

0<n≤100
进阶：时间复杂度 
O(n)


# 解法 滑动窗口

```
func dealSum(start, end int) (sum int, res []int) {

	for i := start; i <= end; i++ {
		sum += i
		res = append(res, i)
	}

	return
}

func FindContinuousSequence(sum int) [][]int {
	// write code here
	var ans [][]int
	if sum == 0 {
		return ans
	}
	i, j := 1, 2

	for i < sum && j < sum {
		s, tmp := dealSum(i, j)

		if s < sum {
			j++
		}

		if s > sum {
			i++
		}

		if s == sum {
			ans = append(ans, tmp)
			i++
			j = i + 1
		}
	}

	return ans
}
```