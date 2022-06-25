package main

func moveZeroes(nums []int) {

	// Start inserting zeroes to the end since we will move them to the end
	currIndexToInsertZero := len(nums) - 1

	for i, v := range nums {
		if v == 0 {
			temp := nums[currIndexToInsertZero]
			nums[currIndexToInsertZero] = 0
			nums[i] = temp
			currIndexToInsertZero--
		}
	}
}
