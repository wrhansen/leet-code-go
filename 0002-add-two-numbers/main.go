package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	array := []int{}
	for {
		array = append(array, node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return fmt.Sprintf("%v", array)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	remainder := 0
	listNode := &ListNode{}
	parentNode := listNode

	for {
		val := l1.Val + l2.Val + remainder
		if val >= 10 {
			val = val - 10
			remainder = 1
		} else {
			remainder = 0
		}
		listNode.Val = val

		if l1.Next == nil && l2.Next == nil {
			if remainder > 0 {
				listNode.Next = &ListNode{Val: remainder}
			}
			break
		}

		listNode.Next = &ListNode{}
		listNode = listNode.Next

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}

	}

	return parentNode
}
