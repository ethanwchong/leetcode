package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	/*var first int
	var second int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			first = nums[i]
			second = nums[i+1]
			break
		}
	}
	if first == 0 && second == 0 {
		return false
	}
	for i := 0; i < len(nums); i++ {

		if nums[i] <= first {
			first = nums[i]
			continue
		}
		if nums[i] <= second && nums[i] > first {
			second = nums[i]
			continue
		}
		if nums[i] > first && nums[i] > second {
			return true
		}
	}
	return false
	*/
	first := math.MaxInt
	second := math.MaxInt
	for _, n := range nums {
		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 3, 2, 5, 3}))    //false
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))    //true
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))    //false
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6})) //true
	fmt.Println(increasingTriplet([]int{1, 2, 4, 5, 3}))    //true
}

/*
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.



Example 1:

Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
Example 2:

Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
Example 3:

Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (3, 4, 5) is valid because nums[3] == 0 < nums[4] == 4 < nums[5] == 6.
*/
