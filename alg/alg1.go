package alg

import (
	"fmt"
)

//快速排序算法
func Fast(nums []int, begin, end int) {
	if begin < end {
		var i, j, tmp int
		i = begin
		j = end
		for i < j {
			tmp = nums[i]
			for i < j && nums[i] <= tmp {
				i++
			}
			for i < j && nums[j] >= tmp {
				j--
			}
			//exchange
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
			fmt.Println(nums, i, j, tmp)
		}
		nums[i] = tmp
		Fast(nums, begin, i-1)
		Fast(nums, i+1, end)
	}
}
