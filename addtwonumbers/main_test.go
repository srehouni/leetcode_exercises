package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers_with_one_item(t *testing.T) {
	list1 := ListNode{1, nil}
	list2 := ListNode{2, nil}

	result := addTwoNumbers(&list1, &list2)
	resultSlice := result.toSlice()

	if !reflect.DeepEqual(resultSlice, []int{3}) {
		t.Errorf("Expected [3], got [%d]", resultSlice)
	}
}

func TestAddTwoNumbers_with_one_item_and_result_zero(t *testing.T) {
	list1 := ListNode{0, nil}
	list2 := ListNode{0, nil}

	result := addTwoNumbers(&list1, &list2)
	resultSlice := result.toSlice()

	if !reflect.DeepEqual(resultSlice, []int{0}) {
		t.Errorf("Expected [3], got [%d]", resultSlice)
	}
}

func TestAddTwoNumbers_with_three_items(t *testing.T) {
	list1 := generateListFromSlice([]int{2, 4, 3})
	list2 := generateListFromSlice([]int{5, 6, 4})

	result := addTwoNumbers(list1, list2).toSlice()

	if !reflect.DeepEqual(result, []int{7, 0, 8}) {
		t.Errorf("Expected [7, 0 ,8], got %d", result)
	}
}

func TestAddTwoNumbers_with_four_items_and_result_with_four_items(t *testing.T) {
	list1 := generateListFromSlice([]int{1, 2, 3, 4})
	list2 := generateListFromSlice([]int{1, 2, 5, 0})

	result := addTwoNumbers(list1, list2).toSlice()

	if !reflect.DeepEqual(result, []int{2, 4, 8, 4}) {
		t.Errorf("Expected [2, 4, 8, 4], got %d", result)
	}
}

func TestAddTwoNumbers_with_four_items_and_result_with_five_items(t *testing.T) {
	list1 := generateListFromSlice([]int{6, 2, 7, 4})
	list2 := generateListFromSlice([]int{5, 2, 5, 5})

	result := addTwoNumbers(list1, list2).toSlice()

	if !reflect.DeepEqual(result, []int{1, 5, 2, 0, 1}) {
		t.Errorf("Expected [1, 5, 2, 0, 1], got %d", result)
	}
}

func TestAddTwoNumbers_with_four_items_and_two_items(t *testing.T) {
	list1 := generateListFromSlice([]int{6, 2, 7, 4})
	list2 := generateListFromSlice([]int{5, 2})

	result := addTwoNumbers(list1, list2).toSlice()

	if !reflect.DeepEqual(result, []int{1, 5, 7, 4}) {
		t.Errorf("Expected [1, 5, 7, 4], got %d", result)
	}
}

func TestAddTwoNumbers_with_seven_items_and_four_items(t *testing.T) {
	list1 := generateListFromSlice([]int{9, 9, 9, 9, 9, 9, 9})
	list2 := generateListFromSlice([]int{9, 9, 9, 9})

	result := addTwoNumbers(list1, list2).toSlice()

	if !reflect.DeepEqual(result, []int{8, 9, 9, 9, 0, 0, 0, 1}) {
		t.Errorf("Expected [8,9,9,9,0,0,0,1], got %d", result)
	}
}
