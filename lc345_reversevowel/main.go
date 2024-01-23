package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	runeSlc := make([]rune, len(s))
	for i, r := range s {
		runeSlc[i] = r
	}

	vowel := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	low, high := 0, len(runeSlc)-1
	for low < high {
		if !vowel[runeSlc[high]] {
			high--
			continue
		}
		if !vowel[runeSlc[low]] {
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
