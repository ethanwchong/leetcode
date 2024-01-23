package main

import "fmt"

func moveZeroes(nums []int) {
	l := 0
	for r := 0; r < len(nums); r++ {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
	}
	for ; l < len(nums); l++ {
		nums[l] = 0
	}
	fmt.Println(nums)

}

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
}
