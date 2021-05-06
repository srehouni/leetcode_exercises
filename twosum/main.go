package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	for index, value := range nums {
		found, resultIndex := checkSum(value, index, nums, target)
		if found {
			return []int{index, resultIndex}
		}
	}

	return []int{}
}

func checkSum(num int, index int, nums []int, target int) (bool, int) {
	for i := index + 1; i < len(nums); i++ {
		if num+nums[i] == target {
			return true, i
		}
	}

	return false, -1
}
