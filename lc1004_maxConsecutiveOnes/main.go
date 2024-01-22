package main

import "fmt"

func longestOnes(nums []int, k int) int {
	maxLen := 0
	zeroes := 0

	for startWindow, endWindow := 0, 0; endWindow < len(nums); endWindow++ {
		if nums[endWindow] == 0 {
			zeroes++
		}

		if zeroes > k {
			if nums[startWindow] == 0 {
				zeroes--
			}
			startWindow++
		}

		windowLen := endWindow - startWindow + 1
		if maxLen < windowLen {
			maxLen = windowLen
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))                         //6
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)) //10
}

/*
Intuition
Translation:
Find the longest subarray with at most K zeros.


Explanation
For each A[j], try to find the longest subarray.
If A[i] ~ A[j] has zeros <= K, we continue to increment j.
If A[i] ~ A[j] has zeros > K, we increment i (as well as j).

*/

/*
Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

Example 1:
Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
Output: 6
Explanation: [1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.

Example 2:
Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
Output: 10
Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.


Constraints:

1 <= nums.length <= 105
nums[i] is either 0 or 1.
0 <= k <= nums.length
*/
