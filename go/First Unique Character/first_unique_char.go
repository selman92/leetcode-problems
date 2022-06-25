package main

import "fmt"

func firstUniqChar(s string) int {
	charCounts := make(map[rune]int)

	for i, v := range s {
		_, exists := charCounts[v]

		if exists {
			charCounts[v] = -1
		} else {
			charCounts[v] = i
		}
	}

	uniqueCharIndex := -1

	for _, v := range charCounts {
		// If the value is -1, this is not a unique char
		if v == -1 {
			continue
		}

		if uniqueCharIndex == -1 {
			uniqueCharIndex = v
		} else if v < uniqueCharIndex {
			uniqueCharIndex = v
		}
	}

	return uniqueCharIndex
}

func main() {
	fmt.Println(firstUniqChar("abbbbacddex"))
}
