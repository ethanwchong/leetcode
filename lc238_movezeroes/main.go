package main

import "fmt"

func moveZeroes(nums []int) {
	var outer = 0
	for outer < len(nums) {
		var p int = 0
		for p != len(nums)-1 {
			if nums[p] == 0 {
				nums[p], nums[p+1] = nums[p+1], nums[p]
			}
			p++
		}
		fmt.Println(nums)
		outer++
	}
}

func moveZeroesAlt(nums []int) {
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[nonZeroIndex] = nums[nonZeroIndex], nums[i]
			nonZeroIndex++
		}
		fmt.Println(nums)
	}
}

func main() {
	//moveZeroes([]int{0, 1, 0, 3, 12})
	//moveZeroes([]int{0})
	//moveZeroes([]int{1, 0, 6, 0, 9, 0, 10, 23})

	moveZeroesAlt([]int{0, 1, 0, 3, 12})
	moveZeroesAlt([]int{0})
	moveZeroesAlt([]int{1, 0, 6, 0, 9, 0, 10, 23})
}

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.



Example 1:

Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]

Example 2:

Input: nums = [0]
Output: [0]
*/
