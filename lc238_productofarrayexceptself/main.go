package main

import "fmt"

func productExceptSelf(nums []int) []int {
	retSlc := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		var sum int = 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				sum = sum * nums[j]
			}
		}
		retSlc = append(retSlc, sum)
	}
	return retSlc
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

*/
