package main

import "fmt"

// time complexity - 2^n
// space complexity - 2^n
func printSubsequence(nums []int, index int, current []int) [][]int {
	if index == len(nums) {
		return [][]int{append([]int{}, current...)}
	}

	// pick the current element
	withCurrent := printSubsequence(nums, index+1, append(current, nums[index]))

	// don't pick the current element
	withoutCurremt := printSubsequence(nums, index+1, current)

	return append(withCurrent, withoutCurremt...)
}

func subsets(nums []int) {
	result := printSubsequence(nums, 0, []int{})
	fmt.Println(result)
}
