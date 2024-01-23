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

// INPUT	[a		b	c	d	]
// LEFT		[1		a	ab	abc	]
// RIGHT	[bcd	cd	d	1	]
// ANS		[bcd	acd	abd	abc ]

func productExceptSelf2(nums []int) []int {
	l := len(nums)
	left := make([]int, l)
	right := make([]int, l)
	ans := make([]int, l)
	left[0] = 1
	right[l-1] = 1
	for i := 1; i < l; i++ {
		left[i] = nums[i-1] * left[i-1]
		j := len(nums) - i - 1
		right[j] = nums[j+1] * right[j+1]
	}
	for i := 0; i < l; i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	fmt.Println(productExceptSelf2([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf2([]int{-1, 1, 0, -3, 3}))
}

/*
238. Product of Array Except Self
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
