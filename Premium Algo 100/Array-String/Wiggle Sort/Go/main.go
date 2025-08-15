package main

import "fmt"

// Wiggle Sort
// Problem: Given an unsorted array nums, reorder it in-place so that nums[0] <= nums[1] >= nums[2] <= nums[3]....

func wiggleSort(nums []int) {
	// TODO: Implement your solution here
}

func main() {
	// Test cases
	testCases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{3, 5, 2, 1, 6, 4},
			expected: []int{3, 5, 1, 6, 2, 4},
		},
		// Add more test cases here
	}

	for i, tc := range testCases {
		input := make([]int, len(tc.nums))
		copy(input, tc.nums)
		wiggleSort(input)
		fmt.Printf("Test case %d: Input %v, Result %v\n", i+1, tc.nums, input)
	}
}
