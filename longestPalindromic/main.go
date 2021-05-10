package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abaxrt"))
}

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var start, end int

	for i := 0; i < len(s); i++ {
		len1 := palindromeLen(s, i, i)
		len2 := palindromeLen(s, i, i+1)

		lenMax := max(len1, len2)

		if lenMax > end-start {
			start = i - ((lenMax - 1) / 2)
			end = i + lenMax/2
		}
	}

	return s[start : end+1]
}

func palindromeLen(s string, i int, j int) int {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return j - i - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
