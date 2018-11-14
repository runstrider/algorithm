package main

import (
	"fmt"
)

func fast_sort(s []int, left, right int) {
	if len(s) <= 1 {
		return
	}
	if left >= right {
		return
	}

	i := left
	j := right
	tmp := s[i]
	for i < j {
		fmt.Println(i, j, len(s), s[j], tmp)
		for s[j] >= tmp {
			j--
		}
		for s[i] <= tmp {
			i++
		}

		if i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}

	}
}

func main() {
	///////////0, 1, 2, 3, 4,   5, 6, 7, 8, 9, 10,11,12,13
	a := []int{3, 2, 0, 4, 33, 22, 5, 7, 8, 0, 1, 4, 3, 2}
	//a := []int{3,9,8,7,6,5,4}
	fmt.Println("ORIGINAL:", a)

	fastsort(a)
	fmt.Println("Fast Result:", a)
	target := 5
	fmt.Println("Below is twosum!")
	fmt.Println(twosum(a, target))

	fmt.Println("Below is threesum!")
	fmt.Println(threesum(a, target))

	fmt.Println("Below is foursum!")
	fmt.Println(foursum(a, target))

	//最大的int数
	//fmt.Println(int(^uint(0)/2))

	fmt.Println("Closet threesum")
	fmt.Println(threecloset(a, target))

	//翻转整数
	fmt.Println("Reverse Integer")
	fmt.Println(reverseInteger(-123))

	//最长回文子串长度
	s := "asdfghjklasdfghjklpoiuytrewqedfcvuj"
	fmt.Println("最长回文子序列的长度")
	fmt.Println(longestPalindromeSubsequenceLength(s))

	fmt.Println("回文子序列的个数")
	fmt.Println(palindromeSubsequenceCount(s))

	fmt.Println("最长回文子串")
	fmt.Println(longPalindrome(s))

	fmt.Println("两数相加twosum")
	fmt.Println(addtwonumber([]int{2, 4, 3}, []int{5, 6, 4}))

	fmt.Println("打劫house")
	fmt.Println(houserobber(a))

	fmt.Println("房屋打劫2")
	fmt.Println(houserobber2(a))

	nums := [][]int{
		{1,0,1,0,0},
		{1,0,1,1,1},
		{1,1,1,1,1},
		{1,0,0,1,0}}
	fmt.Println("maximalsquare")
	fmt.Println(maximalsquare(nums))
}

func quic_wiki(data []int) {
	if len(data) <= 1 {
		return
	}
	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	//data[head] = mid
	quic_wiki(data[:head])
	quic_wiki(data[head+1:])

	return
}

func max(i, j int) int {
	if i < j {
		return j
	} else {
		return i
	}
}

//house rob
func dp_house(nums []int) int {
	dp := make([]int, len(nums))
	dp[1] = nums[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]

}

//exchange
func exchange(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

//bubble
func bubble(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				exchange(&nums[j], &nums[j+1])
			}
		}
	}
}

//fastsort
func fastsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	head := 0
	tail := len(nums) - 1
	tmp := nums[0]
	for i := 1; i <= tail; {
		if nums[i] > tmp {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			nums[i], nums[head] = nums[head], nums[i]
			head++
			i++
		}
	}
	fastsort(nums[:head])
	fastsort(nums[head+1:])
	return
}

//twosum + fastsort
func twosum(nums []int, target int) interface{} {

	type res struct {
		left  int
		right int
	}
	//快速排序
	fastsort(nums)
	left := 0
	right := len(nums) - 1
	var result []res
	for left < right {
		tmp := nums[left] + nums[right]
		if tmp == target {
			result = append(result, res{left, right})
			left++
			for left < len(nums) && nums[left] == nums[left-1] {
				//result = append(result, res{left, right})
				left++
			}
			right--
			for right > 0 && nums[right] == nums[right+1] {
				//result = append(result, res{left, right})
				right--
			}

		} else if tmp > target {
			right--
		} else {
			left++
		}
	}
	return result
}

//3sum
func threesum(nums []int, target int) interface{} {
	fastsort(nums)
	tail := len(nums) - 1
	type res struct {
		i int
		j int
		k int
	}

	var result []res

	for i := 0; i < tail-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//只需往后遍历即可,这里就是twosum
		//twosum(nums[i+1:], target-nums[i])
		left := i + 1
		right := tail
		des := target - nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			if sum == des {
				result = append(result, res{i, left, right})
				right--
				left++

				for left < tail && nums[left] == nums[left-1] {
					left++
				}
				for right > 0 && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > des {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

//3closet
func threecloset(nums []int, target int) interface{} {
	fastsort(nums)
	var closet int
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		min := int(^uint(0) / 2)
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < min {
				min = abs(sum - target)
				closet = sum
			}

			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return sum
			}
		}

	}
	return closet
}

//4sum其实
func foursum(nums []int, target int) interface{} {
	type ret struct {
		i     int
		j     int
		left  int
		right int
	}
	fastsort(nums)
	var result []ret
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			left := j + 1
			right := len(nums) - 1

			for left < right {
				tmp := target - nums[i] - nums[j]
				if nums[left]+nums[right] == tmp {
					result = append(result, ret{i, j, left, right})
					left++
					right--
					for left < len(nums) && nums[left] == nums[left-1] {
						left++
					}
					for right > 0 && nums[right] == nums[right+1] {
						right--
					}
				} else if nums[left]+nums[right] < tmp {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

//reverseInteger
func reverseInteger(num int) int {
	mod := num
	var result int
	var res int
	for mod != 0 {
		tmp := result
		result = result*10 + res
		//需要判断溢出
		if result/10 != tmp {
			return 0
		}
		res = mod % 10
		mod = mod / 10
	}
	if res != 0 {
		result = result*10 + res
	}
	return result
}

//最长回文子序列的最大长度，思想动态规划
// longestPalindromeSubsequenceLength
//经过思考，动态规划思想必须先算小的可能所有组合，然后再算大的，一步一步到最后，自己写的有错误
//一定要先先写状态转换方程
//状态转移方程：   dp[i][j]=dp[i+1][j-1] + 2  if（str[i]==str[j]）
//               dp[i][j]=max(dp[i+1][j],dp[i][j-1])  if （str[i]!=str[j]）
//               计算dp[i][j]时需要计算dp[i+1][*]或dp[*][j-1]，因此i应该从大到小，即递减；j应该从小到大，即递增。
func longestPalindromeSubsequenceLength(str string) int {
	length := len(str)
	fmt.Println("Oridinal String", str)
	var dp [][]int
	dp = make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}
	//这步需要好好想想，其实从后往前做的
	for j := 0; j < length; j++ {
		dp[j][j] = 1
		for i := j - 1; i >= 0; i-- {
			if str[i] == str[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	//fmt.Println(dp)
	return dp[0][length-1]
}

//最长回文子序列个数
func palindromeSubsequenceCount(str string) int {
	//return palindromeSubsequenceCount2(str)
	var dp [][]int
	length := len(str)
	dp = make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}

	for i := length - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j <= length-1; j++ {
			if str[i] != str[j] {
				dp[i][j] = dp[i][j-1] + dp[i+1][j] - dp[i+1][j-1]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i+1][j] - dp[i+1][j-1] + dp[i+1][j-1] + 1
			}
		}
	}

	return dp[0][length-1]
}

//最长回文子序列个数2,同上，只是遍历方向不一样
func palindromeSubsequenceCount2(str string) int {
	length := len(str)
	var dp [][]int
	dp = make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}

	for j := 0; j < length; j++ {
		dp[j][j] = 1
		for i := j - 1; i >= 0; i-- {
			dp[i][j] = dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1]
			if str[i] == str[j] {
				dp[i][j] += 1 + dp[i+1][j-1]
			}
		}
	}
	return dp[0][length-1]
}

//最长回文子串，also动态规划
func longPalindrome(str string) string {
	if str == "" {
		return ""
	}
	length := len(str)
	var dp [][]int
	dp = make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}

	var head, tail, maxlength int

	maxlength = tail - head + 1
	/*
		for i := 0; i < length-1; i++ {
			if str[i] == str[i+1] {
				dp[i][i+1] = 1
				head = i
				tail = i + 1
				maxlength = 2
			}
		}
	*/
	for i := length - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < length; j++ {
			if str[i] == str[j] {
				dp[i][j] = dp[i+1][j-1]
				if dp[i][j] == 1 && (j-i+1 > maxlength) {
					maxlength = j - i + 1
					head = i
					tail = j
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	/*
		for k := 2; k < length; k++ {
			for i := 0; i < length; i++ {
				j := k + i
				if j > length-1 {
					break
				}
				if str[i] == str[j] {
					dp[i][j] = dp[i+1][j-1]
					if dp[i][j] == 1 && k+1 > maxlength {
						maxlength = k + 1
						head = i
						tail = j
					}
				} else {
					dp[i][j] = 0
				}
			}
		}
	*/
	//fmt.Println( dp)
	return str[head : tail+1]
}

//add two number
//2->4->3 + 5->6->4 = 7->0->8 = 807
func addtwonumber(nums1, nums2 []int) []int {

	var result []int
	var sum int
	var carry int
	if len(nums1) < len(nums2) {
		tmp := nums2
		nums2 = nums1
		nums1 = tmp
	}

	for i := 0; i < len(nums1); i++ {
		if i < len(nums2) {
			sum = (nums1[i] + nums2[i] + carry) % 10
			result = append(result, sum%10)
			carry = (nums1[i] + nums2[i]) / 10
		}
	}

	for i := len(nums2); i < len(nums1); i++ {
		sum = (nums1[i] + nums2[i] + carry) % 10
		result = append(result, sum%10)
		carry = (nums1[i] + nums2[i]) / 10
	}
	if carry != 0 {
		result = append(result, carry)
	}
	var res []int
	for i := len(result) - 1; i >= 0; i-- {
		res = append(res, result[i])
	}
	return res
}

//房屋打劫, also动态规划
func houserobber(property []int) int {
	if len(property) == 0 {
		return 0
	}

	if len(property) == 1{
		return property[0]
	}
	var dp []int
	dp = make([]int, len(property))
	dp[0] = property[0]
	dp[1] = max(property[0], property[1])
	for i := 2; i < len(property); i++ {
		dp[i] = max(dp[i-1], dp[i-2] + property[i])
	}
	return dp[len(property)-1]
}

//房屋打劫，houserobber2;
//问题，将房屋连成一体成为一个环，相邻两个房屋不能同时抢；
// 则问题分析：即0和n-1不能同时抢，则max(0~n-2, 1~n-1)即为答案
func houserobber2(property []int) int{
	if len(property) == 0 {
		return 0
	}

	return  max(houserobber(property[1:]), houserobber(property[:len(property)-1]))
}

func min(num1, num2 int) int{
	if num1 < num2 {
		return num1
	}else {
		return  num2
	}
}

//Maximal Square:矩形里最大的正方形
func maximalsquare(matrix [][]int) int{
	var dp [][]int

	if len(matrix) == 1{
		return matrix[0][0]
	}
	lengthx := len(matrix)
	lengthy := len(matrix[0])
	dp = make([][]int, lengthx)
	for i:=0;i<lengthx;i++{
		dp[i]= make([]int, lengthy)
	}

	var maxlength int
	//初始化不能得到的状态值
	for i:=0;i<lengthx;i++{
		dp[i][0] = matrix[i][0]
		if dp[i][0] == 1{
			maxlength = 1
		}
	}
	for i:=0;i<lengthy;i++{
		dp[0][i] = matrix[0][i]
		if dp[0][i] == 1{
			maxlength = 1
		}
	}

	for i:=1;i<lengthx;i++{
		for j:=1;j<lengthy;j++{
			if matrix[i][j] == 1 {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxlength{
					maxlength = dp[i][j]
				}
			}else{
				dp[i][j] = 0
			}
		}
	}
	return maxlength * maxlength
}
