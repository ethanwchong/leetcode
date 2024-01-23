package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	var i, j int = 0, 0
	w1S := strings.Split(word1, "")
	w2S := strings.Split(word2, "")
	var mS []string = make([]string, 0)
	for i < len(w1S) && j < len(w2S) {
		mS = append(mS, w1S[i], w2S[j])
		i++
		j++
	}
	for i < len(w1S) {
		mS = append(mS, w1S[i])
		i++
	}
	for j < len(w2S) {
		mS = append(mS, w2S[j])
		j++
	}

	return strings.Join(mS, "")
}

func mergeAlternately2(word1 string, word2 string) string {
	var i, j int = 0, 0
	stringBuf := ""
	for i < len(word1) && j < len(word2) {
		stringBuf += string(word1[i]) + string(word2[i])
		i++
		j++
	}
	for i < len(word1) {
		stringBuf += string(word1[i])
		i++
	}
	for j < len(word2) {
		stringBuf += string(word2[i])
		j++
	}

	return stringBuf
}

func main() {
	fmt.Println(mergeAlternately("ACE", "BD"))
	fmt.Println(mergeAlternately("MONKEY", "BOY"))
	fmt.Println(mergeAlternately("BOY", "GIRL"))
	fmt.Println(mergeAlternately2("ACE", "BD"))
	fmt.Println(mergeAlternately2("MONKEY", "BOY"))
	fmt.Println(mergeAlternately2("BOY", "GIRL"))

}

/*
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.



Example 1:

Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
Example 2:

Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: Notice that as word2 is longer, "rs" is appended to the end.
word1:  a   b
word2:    p   q   r   s
merged: a p b q   r   s
Example 3:

Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: Notice that as word1 is longer, "cd" is appended to the end.
word1:  a   b   c   d
word2:    p   q
merged: a p b q c   d
*/
