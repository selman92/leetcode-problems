package main

import "fmt"

func rotate(nums []int, k int) {

	numsCopy := make([]int, len(nums))

	copy(numsCopy, nums)
	for i, v := range numsCopy {
		newIndex := k + i

		for newIndex >= len(nums) {
			newIndex = newIndex - len(nums)
		}

		nums[newIndex] = v
	}
}

func main() {
	arr := []int{1, 3, 4, 2, 8}

	rotate(arr, 2)

	fmt.Println(arr)
}
