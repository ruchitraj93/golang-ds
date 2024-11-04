package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := findCombinationWithGivenSum(0, candidates, []int{}, target, 0)
	fmt.Println(result)
	return result
}

func findCombinationWithGivenSum(index int, arr []int, subseq []int, targetSum int, currentSum int) [][]int {

	if index == len(arr) {
		return [][]int{}
	}

	if currentSum == targetSum {
		return [][]int{append([]int{}, subseq...)}
	}

	if currentSum > targetSum {
		return [][]int{}
	}

	// pick the current element
	pick := findCombinationWithGivenSum(index, arr, append(subseq, arr[index]), targetSum, currentSum+arr[index])

	// do not pick the current element
	not_pick := findCombinationWithGivenSum(index+1, arr, subseq, targetSum, currentSum)

	return append(pick, not_pick...)
}
