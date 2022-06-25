package main

import "fmt"

func strStr(haystack string, needle string) int {

	// used to keep track of the index of haystack
	currHaystackIndex := 0

	// used to check whether we currently have a partial match
	matchInProgress := false

	// the index of the needle in haystack
	needleIndex := -1

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[currHaystackIndex] {
			if !matchInProgress {
				needleIndex = i
			}

			matchInProgress = true
			currHaystackIndex++

			// if we reached to the end of needle, it means we found it.
			if currHaystackIndex == len(needle) {
				return needleIndex
			}
		} else {
			if matchInProgress {
				i = needleIndex
			}

			matchInProgress = false
			currHaystackIndex = 0
			needleIndex = -1
		}
	}

	return -1
}

func main() {
	fmt.Print(strStr("mississippi", "pi"))
}
