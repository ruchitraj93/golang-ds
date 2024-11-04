package main

import "fmt"

func findFirstSubsequence(index int, arr []int, subseq []int, targetSum int, currentSum int) bool {
	if index == len(arr) {
		if currentSum == targetSum {
			fmt.Println(subseq)
			return true
		}
		return false
	}

	pick := findFirstSubsequence(index+1, arr, append(subseq, arr[index]), targetSum, currentSum+arr[index])

	if pick {
		return true
	}

	not_pick := findFirstSubsequence(index+1, arr, subseq, targetSum, currentSum)

	return not_pick
}

func printFirstSubsequenceWithSumK(arr []int, targetSum int) {
	findFirstSubsequence(0, arr, []int{}, targetSum, 0)
}
