package main

import "fmt"

func twosum1(num []int, target int) []int {
	for i := 0; i < len(num)-1; i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}

}

func twosum2(num []int, target int) []int {
	m := map[int]int{} //key = array item's value; value= array item's index
	for i := 0; i < len(num); i++ {
		if matchIndex, isPresent := m[target-num[i]]; isPresent {
			return []int{matchIndex, i}
		}
		m[num[i]] = i
	}
	return []int{}

}

func main() {
	var nums []int = []int{2, 3, 8, 7}
	var target = 10
	fmt.Printf("the value is %d", twosum1(nums, target))
	fmt.Printf("\nthe value is %d", twosum2(nums, target))

}

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/
