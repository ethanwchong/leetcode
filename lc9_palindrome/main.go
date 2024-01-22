package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindromeBad(num int) bool {
	numStr := strconv.Itoa(num)
	numSlice := strings.Split(numStr, "")
	length := len(numSlice) / 2
	firstHalf := numSlice[:length]
	fmt.Println(firstHalf)
	secondHalf := numSlice[length:]
	fmt.Println(secondHalf)
	var secondHalfSize = len(secondHalf)
	for _, v := range firstHalf {
		secondHalfSize = secondHalfSize - 1
		if v != secondHalf[secondHalfSize] {
			return false
		}
	}
	return true
}

func reverse(str string) string {
	//convert string to slice
	strSlc := strings.Split(str, "")
	lastIndex := len(strSlc)
	for i := 0; i < len(strSlc)/2; i++ {
		lastIndex = lastIndex - 1
		strSlc[i], strSlc[lastIndex] = strSlc[lastIndex], strSlc[i]
	}
	return strings.Join(strSlc, "")
}

func isPalindromeBetter(num int) bool {
	//convert int to string
	numStr := strconv.Itoa(num)
	lastIndex := len(numStr)
	for i := 0; i < len(numStr)/2; i++ {
		lastIndex = lastIndex - 1
		if numStr[i] != numStr[lastIndex] {
			return false
		}
	}
	return true
}

func main() {
	x := 12355321
	fmt.Println(isPalindromeBad(x))
	fmt.Println(isPalindromeBetter(x))
	fmt.Println(reverse("ABCDEFGH"))
}

/*
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

*/
