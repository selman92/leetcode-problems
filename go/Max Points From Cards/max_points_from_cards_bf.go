package main

import "fmt"

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}

func getSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

func maxScore(cardPoints []int, k int) int {

	if k == len(cardPoints) {
		return getSum(cardPoints)
	}

	max := 0
	score := 0
	for i := 0; i <= k; i++ {
		elementsFromStart := i
		elementsFromEnd := k - i

		//fmt.Println("CARDPOINTS: ", cardPoints)
		fmt.Println("From start: ", i, " From END: ", elementsFromEnd)
		if elementsFromStart == 0 {
			//fmt.Println("Getting sum of: ", cardPoints[len(cardPoints)-elementsFromEnd:])
			score = getSum(cardPoints[len(cardPoints)-elementsFromEnd:])
		} else if elementsFromEnd == 0 {
			//fmt.Println("Getting sum of: ", cardPoints[0:elementsFromStart])
			score = getSum(cardPoints[0:elementsFromStart])
		} else {
			//fmt.Println("Getting sum of: ", cardPoints[0:elementsFromStart], cardPoints[len(cardPoints)-elementsFromEnd:])
			score = getSum(cardPoints[0:elementsFromStart]) + getSum(cardPoints[len(cardPoints)-elementsFromEnd:])
		}

		if score > max {
			max = score
		}
	}

	return max
}
