package main

import "fmt"

func gcdOfStrings(s1 string, s2 string) string {
	if s1 == s2 {
		return s1
	}
	if s1+s2 != s2+s1 {
		return ""
	}
	return s2[:gcd(len(s1), len(s2))]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
}
