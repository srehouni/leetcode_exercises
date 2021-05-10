package main

import (
	"testing"
)

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

	if palindrome != "a" && palindrome != "c" {
		t.Errorf("Expected a or c, got %s", palindrome)
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
