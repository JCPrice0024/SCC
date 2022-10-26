package SCC

import "sort"

// SubtractDups takes in a slice of ints and checks to see if there are any duplicate values.
// If there are it skips over them, otherwise it adds the values and returns an integer value of the result
func SubtractDups(nums []int) int {
	var correctTotal int
	nums = append(nums, 0)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			correctTotal += nums[i]
		}
	}
	return correctTotal
}
