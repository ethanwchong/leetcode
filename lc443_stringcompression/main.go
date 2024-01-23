package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	l, r := 0, 0

	for r < len(chars) {
		j := r + 1
		for ; j < len(chars) && chars[j] == chars[j-1]; j++ {
		}

		chars[l] = chars[r]
		l++
		if j != r+1 {
			for _, number := range strconv.Itoa(j - r) {
				chars[l] = byte(number)
				l++
			}
		}
		r = j
	}

	chars = chars[:l]
	return len(chars)
}

func compress2(chars []byte) int {
	//var writer int = 0
	//var nextChar int = 0
	var charCounter int = 0

	for i := 0; i < len(chars)-1; i++ {
		if chars[i] != chars[i+1] {
			charCounter++
		}
	}
	return (charCounter + 1) * 2

}

func main() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))                               //6
	fmt.Println(compress([]byte{'a'}))                                                             //1
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'})) //4
	fmt.Println(compress([]byte{'a', 'a', 'a', 'b', 'b', 'b', 'c', 'c', 'c'}))                     //6
	fmt.Println("==========")
	fmt.Println(compress2([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}))                               //6
	fmt.Println(compress2([]byte{'a'}))                                                             //1
	fmt.Println(compress2([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'})) //4
	fmt.Println(compress2([]byte{'a', 'a', 'a', 'b', 'b', 'b', 'c', 'c', 'c'}))                     //6
}

/*
443. String Compression

Given an array of characters chars, compress it using the following algorithm:

Begin with an empty string s. For each group of consecutive repeating characters in chars:

If the group's length is 1, append the character to s.
Otherwise, append the character followed by the group's length.
The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

After you are done modifying the input array, return the new length of the array.

You must write an algorithm that uses only constant extra space.



Example 1:

Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
Example 2:

Input: chars = ["a"]
Output: Return 1, and the first character of the input array should be: ["a"]
Explanation: The only group is "a", which remains uncompressed since it's a single character.
Example 3:

Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
Output: Return 4, and the first 4 characters of the input array should be: ["a","b","1","2"].
Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".

*/
