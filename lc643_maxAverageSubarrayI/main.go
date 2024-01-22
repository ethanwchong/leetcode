package main

import "fmt"

// Sum up the first K window.
// Create a pointer P, pointing at the beginning of the next window.
// Replace by iterate the int slice from P until the end of the slice one integer at a time.
//

func findMaxAverage(nums []int, k int) float64 {
	//find the max value of the 1st window
	var max int = 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}

	//Move the window forward by 1 and sum the new window and deduct from
	//the 1st element of the previous window.
	//Check to see if the new sum is bigger than the previous sum.
	var p int = k
	for p < len(nums) {
		newSum := max + nums[p] - nums[p-k]
		if newSum > max {
			max = newSum
			newSum = 0
		}
		p++
	}
	return float64(max) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) //12.75
	fmt.Println(findMaxAverage([]int{5}, 1))
}

/*
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.



Example 1:

Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
Example 2:

Input: nums = [5], k = 1
Output: 5.00000


Constraints:

n == nums.length
1 <= k <= n <= 105
-104 <= nums[i] <= 104
*/
