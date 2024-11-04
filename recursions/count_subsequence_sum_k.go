package main

import "fmt"

func findSubsequenceWithSumK(index int, arr []int, targetSum int, currentSum int) int {
	if index == len(arr) {
		if currentSum == targetSum {
			return 1
		}
		return 0
	}

	// pick the current element
	count_pick := findSubsequenceWithSumK(index+1, arr, targetSum, currentSum+arr[index])

	count_not_pick := findSubsequenceWithSumK(index+1, arr, targetSum, currentSum)

	return count_pick + count_not_pick
}

func countSubsequenceWithSumK(arr []int, targetSum int) {
	result := findSubsequenceWithSumK(0, arr, targetSum, 0)

	fmt.Println(result)
}
