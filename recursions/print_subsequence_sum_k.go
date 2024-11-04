package main

import "fmt"

func findSubsequence(index int, arr []int, subseq []int, targetSum int, currentSum int) {
	if index == len(arr) {
		if targetSum == currentSum {
			fmt.Println(subseq)
		}
		return
	}

	// pick the current element
	findSubsequence(index+1, arr, append(subseq, arr[index]), targetSum, currentSum+arr[index])

	// do not pick the current element
	findSubsequence(index+1, arr, subseq, targetSum, currentSum)
}

func printSubsequencesWithSumK(arr []int, k int) {
	findSubsequence(0, arr, []int{}, k, 0)
}
