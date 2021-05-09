package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(lengthOfLongestSubstring("Hello"))
	fmt.Println(strings.Split("deervdf", "d"))
	fmt.Println(len(strings.Split("dvdf", "d")))
}

func lengthOfLongestSubstring(s string) int {
	comparingFirstSubstring := true
	longestSubstring := ""
	currentSubstring := ""

	for i := 0; i < len(s); i++ {
		if comparingFirstSubstring {
			if !strings.Contains(longestSubstring, string(s[i])) {
				longestSubstring += string(s[i])
			} else {
				comparingFirstSubstring = false
				substrings := strings.Split(longestSubstring, string(s[i]))
				substring := substrings[len(substrings)-1] + string(s[i])
				currentSubstring += substring
			}
		} else {
			if !strings.Contains(string(currentSubstring), string(s[i])) {
				currentSubstring += string(s[i])
			} else {
				substrings := strings.Split(currentSubstring, string(s[i]))
				substring := substrings[len(substrings)-1] + string(s[i])

				if len(longestSubstring) >= len(currentSubstring) {
					currentSubstring = ""
				} else {
					longestSubstring = currentSubstring
					currentSubstring = ""
				}
				currentSubstring = substring
			}
		}
	}

	if len(longestSubstring) >= len(currentSubstring) {
		return len(longestSubstring)
	} else {
		return len(currentSubstring)
	}
}
