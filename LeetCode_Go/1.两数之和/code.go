package main

import (
	"fmt"
)

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
func main() {
	var test = []int{3, 2, 4}
	target := 6
	result := twoSum2(test, target)
	fmt.Println(result)
}

// 暴力法
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

// 使用map哈希查找
// {3, 2, 4} 6
func twoSum2(nums []int, target int) []int {
	result := make([]int, 0)
	maps := make(map[int]int)
	for key, val := range nums {
		_, dp := maps[target-val]
		if dp {
			result = append(result, maps[target-val])
			result = append(result, key)
			break
		}
		maps[val] = key
	}
	return result
}
