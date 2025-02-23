package main

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	input1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	input2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	actual := addTwoNumbers(input1, input2)
	expected := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 8,
			},
		},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("TestCase1 failed! expected: %v, actual: %v\n", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	input1 := &ListNode{Val: 0}
	input2 := &ListNode{Val: 0}

	actual := addTwoNumbers(input1, input2)
	expected := &ListNode{Val: 0}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("TestCase2 failed! expected: %v, actual: %v\n", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	input1 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}

	input2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}

	actual := addTwoNumbers(input1, input2)
	expected := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 1,
								},
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("TestCase3 failed! expected: %v, actual: %v\n", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	input1 := &ListNode{Val: 9}
	input2 := &ListNode{Val: 9}

	actual := addTwoNumbers(input1, input2)
	expected := &ListNode{
		Val:  8,
		Next: &ListNode{Val: 1},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("TestCase4 failed! expected: %v, actual: %v\n", expected, actual)
	}
}
