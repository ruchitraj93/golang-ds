package main

import (
	"fmt"
	"sort"
)

func findSubsetSum(index int, arr []int, subset []int, currentSum int, size int) []int {
	if index == size {
		subset = append(subset, currentSum)
		return subset
	}

	pick := findSubsetSum(index+1, arr, subset, currentSum+arr[index], size)

	not_pick := findSubsetSum(index+1, arr, subset, currentSum, size)

	return append(pick, not_pick...)

}

func subsetSum(arr []int, n int) {
	result := findSubsetSum(0, arr, []int{}, 0, n)
	sort.Ints(result)
	fmt.Println(result)
}
