package main

import (
	"fmt"
)

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
func main() {
	var test = []int{3, 2, 4}
	target := 6
	result := twoSum(test, target)
	fmt.Printf("%v", result)
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 1 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if sum := nums[i] + nums[j]; sum == target {
				var result = []int{i, j}
				return result
			}
		}
	}
	return nil
}
