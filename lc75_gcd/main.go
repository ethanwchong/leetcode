package main

import "fmt"

func gcdOfStrings(s1 string, s2 string) string {
	if s1 == s2 {
		return s1
	}

	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}

	if s1[:len(s2)] != s2 {
		return ""
	}

	return gcdOfStrings(s1[len(s2):], s2)
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("ABC", "ABCABC"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
	fmt.Println(gcdOfStrings("ABCABCABC", "ABCABC"))
	fmt.Println(gcdOfStrings("ABAB", "ABABBC"))

}

/*
For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.



Example 1:

Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Example 2:

Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Example 3:

Input: str1 = "LEET", str2 = "CODE"
Output: ""

*/
