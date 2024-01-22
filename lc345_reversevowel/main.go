package main

import (
	"fmt"
)

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}

func reverseVowels(s string) string {
	//make a slice of rune
	runeSlc := make([]rune, len(s))

	low, high := 0, len(runeSlc)-1

	for i, r := range s {
		runeSlc[i] = r
	}

	for low < high {
		if !isVowel(runeSlc[high]) {
			high--
			continue
		}
		if !isVowel(runeSlc[low]) {
			low++
			continue
		}
		runeSlc[high], runeSlc[low] = runeSlc[low], runeSlc[high]
		high--
		low++
	}

	return string(runeSlc)
}

func main() {
	fmt.Printf("Original: %s; Reversed: %s", "hello", reverseVowels("hello"))         //expected: holle
	fmt.Printf("\nOriginal: %s; Reversed: %s", "leetcode", reverseVowels("leetcode")) //expected: leotcede

}

/*Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.



Example 1:

Input: s = "hello"
Output: "holle"
Example 2:

Input: s = "leetcode"
Output: "leotcede"

*/
