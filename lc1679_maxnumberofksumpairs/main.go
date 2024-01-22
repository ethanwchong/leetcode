package main

import (
	"fmt"
)

func maxOperations(nums []int, k int) int {
	//sort.Ints(nums)
	nums = sort1(nums)
	left := 0
	right := len(nums) - 1
	count := 0

	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			left++
			right--
			count++
		} else if sum < k {
			left++
		} else {
			right--
		}
	}
	return count
}

func sort1(nums []int) []int {
	slcSize := len(nums)
	for i := 0; i < slcSize-1; i++ {
		for j := 1; j < slcSize-1; j++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
		}
	}
	fmt.Println(nums)
	return nums
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))    //2
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6)) //1
}

/*

You are given an integer array nums and an integer k.

In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

Return the maximum number of operations you can perform on the array.



Example 1:

Input: nums = [1,2,3,4], k = 5
Output: 2
Explanation: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
There are no more pairs that sum up to 5, hence a total of 2 operations.
Example 2:

Input: nums = [3,1,3,4,3], k = 6
Output: 1
Explanation: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
There are no more pairs that sum up to 6, hence a total of 1 operation.

*/
