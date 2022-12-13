package main

func longestCommonPrefix(strs []string) string {
	var prefix []uint8
	isCommonPrefix := true
	var curChar uint8
	var i = 0
	var minLen = 0
	for isCommonPrefix {
		if curChar != 0 {
			prefix = append(prefix, curChar)
		}
		curChar = 0
		for _, word := range strs {

			if minLen == 0 && len(word) != 0 {
				minLen = len(word)
			}
			if minLen > len(word) {
				minLen = len(word)
			}

			if curChar == 0 && i < minLen {
				curChar = word[i]
			} else {
				if len(word) == 0 || minLen <= i || word[i] != curChar {
					isCommonPrefix = false
				}
			}
		}
		i++
	}
	return string(prefix)
}

func main() {
	strs := []string{"dog", "racecar", "car"}

	print(longestCommonPrefix(strs))
}
