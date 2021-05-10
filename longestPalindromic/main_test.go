package main

import (
	"testing"
)

func longestPalindromev2(s string) string {
	if len(s) == 0 {
		return s
	}

	var start, end int
	for i := 0; i < len(s); i++ {
		// Case 1
		len1 := expandSubstring(s, i, i)
		// Case 2
		len2 := expandSubstring(s, i, i+1)

		lenMax := max(len1, len2)
		if lenMax > end-start {
			start = i - ((lenMax - 1) / 2)
			end = i + (lenMax / 2)
		}
	}
	return s[start : end+1]
}

func expandSubstring(s string, l, r int) int {
	if len(s) == 0 || l > r {
		return 0
	}

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l -= 1
		r += 1
	}
	return r - l - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindrome(s string) string {
	longest := ""

	if len(s) == 1 {
		return s
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && len(longest) < len(s[i:j+1]) {
				longest = s[i : j+1]
			} else {
				if len(longest) == 0 {
					longest = string(s[i])
				}
			}
		}
	}

	return longest
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func TestLongestPalindrome_one_characters(t *testing.T) {
	s := "a"
	palindrome := longestPalindrome(s)

	if palindrome != "a" {
		t.Errorf("Expected a, got %s", palindrome)
	}
}

func TestLongestPalindrome_one_characters_example1(t *testing.T) {
	s := "ac"
	palindrome := longestPalindrome(s)

	if palindrome != "a" {
		t.Errorf("Expected a, got %s", palindrome)
	}
}

func TestLongestPalindrome_two_same_characters_example1(t *testing.T) {
	s := "bb"
	palindrome := longestPalindrome(s)

	if palindrome != "bb" {
		t.Errorf("Expected bb, got %s", palindrome)
	}
}

func TestLongestPalindrome_three_same_characters_example1(t *testing.T) {
	s := "ccc"
	palindrome := longestPalindrome(s)

	if palindrome != "ccc" {
		t.Errorf("Expected ccc, got %s", palindrome)
	}
}

func TestLongestPalindrome_two_characters(t *testing.T) {
	s := "cbbd"
	palindrome := longestPalindrome(s)

	if palindrome != "bb" {
		t.Errorf("Expected bb, got %s", palindrome)
	}
}

func TestLongestPalindrome_three_characters(t *testing.T) {
	s := "babad"
	palindrome := longestPalindrome(s)

	if palindrome != "bab" && palindrome != "aba" {
		t.Errorf("Expected bab or aba, got %s", palindrome)
	}
}

func TestLongestPalindrome_more_than_one_palindoreme(t *testing.T) {
	s := "aacabdkacaa"
	palindrome := longestPalindrome(s)

	if palindrome != "aca" {
		t.Errorf("Expected aca, got %s", palindrome)
	}
}
