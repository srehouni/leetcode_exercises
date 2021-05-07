package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := generateListFromSlice([]int{1, 2, 9, 2, 4, 6})
	list2 := generateListFromSlice([]int{2, 2, 8, 5})

	result := addTwoNumbers(list1, list2).toSlice()
	fmt.Println(result)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	addToNext := 0
	result := &ListNode{}
	aux := result
	node1 := l1
	node2 := l2

	for node1 != nil || node2 != nil || addToNext > 0 {
		node1Val := 0
		node2Val := 0

		if node1 != nil {
			node1Val = node1.Val
			node1 = node1.Next
		}

		if node2 != nil {
			node2Val = node2.Val
			node2 = node2.Next
		}

		sum = (node1Val + node2Val + addToNext) % 10
		addToNext = int((node1Val + node2Val + addToNext) / 10)

		aux.Next = &ListNode{sum, nil}
		aux = aux.Next
	}

	return result.Next
}

func (l *ListNode) toSlice() []int {
	var slice = []int{}
	node := l

	for node != nil {
		slice = append(slice, node.Val)
		node = node.Next
	}
	return slice
}

func generateListFromSlice(items []int) *ListNode {
	result := &ListNode{}
	aux := result
	for _, value := range items {
		aux.Next = &ListNode{value, nil}
		aux = aux.Next
	}

	return result.Next
}
