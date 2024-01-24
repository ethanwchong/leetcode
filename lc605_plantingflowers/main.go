package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if i == 0 && flowerbed[i+1] != 1 {
			flowerbed[i] = 1
			n--
			continue
		}
		if flowerbed[i] != 1 && flowerbed[i-1] != 1 && flowerbed[i+1] != 1 {
			flowerbed[i] = 1
			n--
			continue
		}
	}
	return n == 0
}

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))          //true
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))          //false
	fmt.Println(canPlaceFlowers([]int{1, 0, 1, 0, 1}, 2))          //false
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 1, 0}, 1)) //true
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 1, 0}, 2)) //false
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0, 1, 0, 1, 0}, 2)) //true
	fmt.Println(canPlaceFlowers([]int{0, 0, 0, 0, 1, 0, 1, 0}, 3)) //false
}

/*
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.



Example 1:

Input: flowerbed = [1,0,0,0,1], n = 1
Output: true
Example 2:

Input: flowerbed = [1,0,0,0,1], n = 2
Output: false

*/
