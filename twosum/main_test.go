package main

import (
	"reflect"
	"testing"
)

func TestCheckSum(t *testing.T) {
	array := []int{1, 2, 3}
	found, index := checkSum(1, 0, array, 3)
	if found && array[index] == 3 {
		t.Errorf("Expected index = 1, got index = %d", index)
	}
}

func TestTwoSum_example1_when_succeeds(t *testing.T) {
	array := []int{1, 2, 3}
	target := 3
	result := twoSum(array, target)

	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("Expected [0, 1], got [%d, %d]", result[0], result[1])
	}
}

func TestTwoSum_example2_when_succeeds(t *testing.T) {
	array := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(array, target)

	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("Expected [0, 1], got [%d, %d]", result[0], result[1])
	}
}

func TestTwoSum_example3_when_succeeds(t *testing.T) {
	array := []int{3, 2, 4}
	target := 6
	result := twoSum(array, target)

	if !reflect.DeepEqual(result, []int{1, 2}) {
		t.Errorf("Expected [1, 2], got [%d, %d]", result[0], result[1])
	}
}

func TestTwoSum_example4_when_succeeds(t *testing.T) {
	array := []int{3, 3}
	target := 6
	result := twoSum(array, target)

	if !reflect.DeepEqual(result, []int{0, 1}) {
		t.Errorf("Expected [0, 1], got [%d, %d]", result[0], result[1])
	}
}

func TestTwoSum_example1_when_fails(t *testing.T) {
	array := []int{1, 2, 3}
	target := 8
	result := twoSum(array, target)

	if !reflect.DeepEqual(result, []int{}) {
		t.Errorf("Expected an empty slice, got %d", result)
	}
}

func TestTwoSum_example2_when_fails(t *testing.T) {
	array := []int{2, 7, 11, 15}
	target := 10
	result := twoSum(array, target)

	if !reflect.DeepEqual(result, []int{}) {
		t.Errorf("Expected an empty slice, got %d", result)
	}
}

func TestTwoSum_example3_when_fails(t *testing.T) {
	array := []int{3, 2, 4}
	target := 1
	result := twoSum(array, target)

	if !reflect.DeepEqual(result, []int{}) {
		t.Errorf("Expected an empty slice, got %d", result)
	}
}

func TestTwoSum_example4_when_fails(t *testing.T) {
	array := []int{3, 3}
	target := 7
	result := twoSum(array, target)

	if !reflect.DeepEqual(result, []int{}) {
		t.Errorf("Expected an empty slice, got %d", result)
	}
}
