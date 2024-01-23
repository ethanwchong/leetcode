package main

import "fmt"

func longestSubarray(nums []int) int {
	left := 0
	right := 0
	zerosCounter := 0
	maxSubArrayWith1s := 0
	for i := 0; i < len(nums); i++ {
		//add 1 to zerosCounter
		if nums[i] == 0 {
			zerosCounter++
		}
		//increase left pointer if zeroCounter > 1 and decrement zeroCounter
		for zerosCounter > 1 {
			if nums[left] == 0 {
				zerosCounter--
			}
			left++
		}
		if right-left > maxSubArrayWith1s {
			maxSubArrayWith1s = right - left
		}
		right++
	}
	return maxSubArrayWith1s
}

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 0, 1}))                //3
	fmt.Println(longestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1})) //5
	fmt.Println(longestSubarray([]int{1, 1, 1}))                   //2
}

/*
Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array. Return 0 if there is no such subarray.



Example 1:

Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
Example 2:

Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].

Example 3:
Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.


Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.

*/
