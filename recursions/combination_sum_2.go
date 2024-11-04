package main

// require revision

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // Sorting helps handle duplicates easily
	var result [][]int
	findCombination(0, target, candidates, []int{}, &result)
	return result
}

func findCombination(index, target int, candidates, current []int, result *[][]int) {
	if target == 0 {
		comb := append([]int{}, current...) // Create a copy of `current`
		*result = append(*result, comb)
		return
	}

	for i := index; i < len(candidates); i++ {
		// Skip duplicate elements to avoid duplicate combinations
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}

		// Stop if the candidate is larger than the remaining target
		if candidates[i] > target {
			break
		}

		findCombination(i+1, target-candidates[i], candidates, append(current, candidates[i]), result)
	}
}
