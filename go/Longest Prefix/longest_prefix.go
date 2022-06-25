package main

func longestCommonPrefix(strs []string) string {
	matchFailed := false
	prefix := make([]rune, 0)

	for k := 0; !matchFailed; k++ {

		if k == len(strs[0]) {
			break
		}

		currChar := strs[0][k]

		for _, v := range strs {
			if k >= len(v) || v[k] != currChar {
				matchFailed = true
				break
			}
		}

		if !matchFailed {
			prefix = append(prefix, rune(currChar))
		}
	}

	return string(prefix)
}
