package main

import (
	"testing"
)

func TestLength_of_longest_substring_with_no_repeated_characters(t *testing.T) {
	s := "hola"

	if lengthOfLongestSubstring(s) != 4 {
		t.Errorf("Expected 4, got %d", lengthOfLongestSubstring(s))
	}
}

func TestLength_of_longest_substring_with_one_repeated_character(t *testing.T) {
	s := "hello"

	if lengthOfLongestSubstring(s) != 3 {
		t.Errorf("Expected 3, got %d", lengthOfLongestSubstring(s))
	}
}

func TestLength_of_longest_substring_with_one_repeated_character_example2(t *testing.T) {
	s := "dvdf"

	if lengthOfLongestSubstring(s) != 3 {
		t.Errorf("Expected 3, got %d", lengthOfLongestSubstring(s))
	}
}

func TestLength_of_longest_substring_with_all_repeated_characters(t *testing.T) {
	s := "bbbbb"

	if lengthOfLongestSubstring(s) != 1 {
		t.Errorf("Expected 1, got %d", lengthOfLongestSubstring(s))
	}
}

func TestLength_of_longest_substring_with_first_substring_as_longest(t *testing.T) {
	s := "abcabcbb"

	if lengthOfLongestSubstring(s) != 3 {
		t.Errorf("Expected 3, got %d", lengthOfLongestSubstring(s))
	}
}

func TestLength_of_longest_substring_with_second_substring_as_longest(t *testing.T) {
	s := "pwwkew"

	if lengthOfLongestSubstring(s) != 3 {
		t.Errorf("Expected 3, got %d", lengthOfLongestSubstring(s))
	}
}

func TestLength_of_longest_substring_with_second_substring_as_longest_second_test(t *testing.T) {
	s := "aab"

	if lengthOfLongestSubstring(s) != 2 {
		t.Errorf("Expected 2, got %d", lengthOfLongestSubstring(s))
	}
}

func TestLength_of_longest_substring_with_empty_string(t *testing.T) {
	s := ""

	if lengthOfLongestSubstring(s) != 0 {
		t.Errorf("Expected 0, got %d", lengthOfLongestSubstring(s))
	}
}
