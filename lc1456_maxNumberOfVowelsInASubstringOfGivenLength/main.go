package main

import (
	"fmt"
	"strings"
)

func isVowel(c string) bool {
	return c == "a" || c == "e" || c == "i" || c == "o" || c == "u" || c == "A" || c == "E" || c == "I" || c == "O" || c == "U"
}

func maxVowelsBad(s string, k int) int {
	max := 0
	slc := strings.Split(s, "")
	for i := 0; i < k; i++ {
		if isVowel(slc[0]) {
			max++
		}
	}
	p := k
	for p < len(slc) {
		newSlc := slc[p-k+1 : p+1]
		tempInt := 0
		for _, v := range newSlc {
			if isVowel(v) {
				tempInt++
			}
		}
		if tempInt > max {
			max = tempInt
		}
		p++
	}

	return max
}

func maxVowels(s string, k int) int {
	vowel := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	//get the max from the 1st window
	max := 0
	for i := 0; i < k; i++ {
		if vowel[s[i]] {
			max++
		}
	}
	if max == k {
		return max
	}
	//slide the window 1 element at a time
	p := k // pointer start at position k
	var evalMax int = max
	for p < len(s) {
		//add the next element to the evalMax
		if vowel[s[p]] {
			evalMax++
		}
		//remove the 1st element from the previous window
		if vowel[s[p-k]] {
			evalMax--
		}
		if evalMax == k {
			return evalMax
		}
		p++
	}
	return max
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3)) //3
	fmt.Println(maxVowels("aeiou", 2))     //2
	fmt.Println(maxVowels("leetcode", 3))  //2

	//fmt.Println(maxVowelsBad("abciiidef", 3)) //3
	//fmt.Println(maxVowelsBad("aeiou", 2))     //2
	//fmt.Println(maxVowelsBad("leetcode", 3))  //2
}

/*
Given a string s and an integer k, return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.

Example 1:

Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.

Example 2:

Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.

Example 3:

Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.


Constraints:

1 <= s.length <= 105
s consists of lowercase English letters.
1 <= k <= s.length
*/
