package main

import (
	"testing"
)

func TestMedian_arrays_with_two_items_and_one_item(t *testing.T) {
	array1 := []int{1, 3}
	array2 := []int{2}
	median := findMedianSortedArrays(array1, array2)

	if median != 2 {
		t.Errorf("Expected 2, got %f", median)
	}
}

func TestMedian_arrays_with_two_items_and_one_item_example2(t *testing.T) {
	array1 := []int{1}
	array2 := []int{2, 3}
	median := findMedianSortedArrays(array1, array2)

	if median != 2 {
		t.Errorf("Expected 2, got %f", median)
	}
}

func TestMedian_arrays_with_two_items_each_example1(t *testing.T) {
	array1 := []int{1, 2}
	array2 := []int{3, 4}
	median := findMedianSortedArrays(array1, array2)

	if median != 2.5 {
		t.Errorf("Expected 2.5, got %f", median)
	}
}

func TestMedian_arrays_with_two_items_each_example2(t *testing.T) {
	array1 := []int{1, 3}
	array2 := []int{2, 7}
	median := findMedianSortedArrays(array1, array2)

	if median != 2.5 {
		t.Errorf("Expected 2.5, got %f", median)
	}
}

func TestMedian_arrays_with_two_zeros_each(t *testing.T) {
	array1 := []int{0, 0}
	array2 := []int{0, 0}
	median := findMedianSortedArrays(array1, array2)

	if median != 0.0 {
		t.Errorf("Expected 0.0, got %f", median)
	}
}

func TestMedian_one_empty_array_example1(t *testing.T) {
	array1 := []int{}
	array2 := []int{1}
	median := findMedianSortedArrays(array1, array2)

	if median != 1.0 {
		t.Errorf("Expected 1.0, got %f", median)
	}
}

func TestMedian_one_empty_array_example2(t *testing.T) {
	array1 := []int{2}
	array2 := []int{}
	median := findMedianSortedArrays(array1, array2)

	if median != 2.0 {
		t.Errorf("Expected 2.0, got %f", median)
	}
}

func TestMedian_one_array_with_negative_numbers(t *testing.T) {
	array1 := []int{3}
	array2 := []int{-2, -1}
	median := findMedianSortedArrays(array1, array2)

	if median != -1.0 {
		t.Errorf("Expected -1.0, got %f", median)
	}
}

func TestMedian_array_with_different_len(t *testing.T) {
	array1 := []int{4, 5}
	array2 := []int{1, 2, 3}
	median := findMedianSortedArrays(array1, array2)

	if median != 3.0 {
		t.Errorf("Expected 3.0, got %f", median)
	}
}

func TestMedian_array_with_different_len_example2(t *testing.T) {
	array1 := []int{4, 5}
	array2 := []int{1, 2, 3, 6, 7}
	median := findMedianSortedArrays(array1, array2)

	if median != 4.0 {
		t.Errorf("Expected 4.0, got %f", median)
	}
}
