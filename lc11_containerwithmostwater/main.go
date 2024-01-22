package main

import "fmt"

func maxArea(height []int) int {
	var maxArea int = 0
	var left int = 0
	var right int = len(height) - 1

	lHeight := 0
	rHeight := 0

	for left < right {
		lHeight = height[left]
		rHeight = height[right]
		if lHeight < rHeight {
			if eval := lHeight * (right - left); maxArea < eval {
				maxArea = eval
			}
			left++
		} else {
			if eval := rHeight * (right - left); maxArea < eval {
				maxArea = eval
			}
			right--
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))  // 49
	fmt.Println(maxArea([]int{1, 1}))                       // 1
	fmt.Println(maxArea([]int{10, 8, 6, 8, 7, 4, 8, 3, 2})) // 49
}

/*
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:

[PICTURE of a Vertical Bar Chart, x is always 1 point apart, heights varies for each bar]

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
Example 2:

Input: height = [1,1]
Output: 1

*/
