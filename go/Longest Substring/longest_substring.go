package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	temp := make([]byte, 0)
	var longestStrLength int = 0
	startIndex := 0
	charMap := make(map[byte]bool, 0)
	for i := 0; i < len(s); i++ {
		_, exists := charMap[s[i]]

		if exists {
			if len(temp) > longestStrLength {
				longestStrLength = len(temp)
			}

			temp = make([]byte, 0)
			charMap = make(map[byte]bool, 0)
			i = startIndex
		} else {
			if len(temp) == 0 {
				startIndex = i
			}

			charMap[s[i]] = true
			temp = append(temp, s[i])
		}
	}

	if len(temp) > longestStrLength {
		longestStrLength = len(temp)
	}

	return longestStrLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
