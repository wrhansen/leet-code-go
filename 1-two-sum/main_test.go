package main

import (
	"reflect"
	"sync"
	"testing"
)

type TestCase struct {
	nums     []int
	target   int
	expected []int
}

func (tc *TestCase) Test(index int, t *testing.T) {
	actual := twoSum(tc.nums, tc.target)
	if !reflect.DeepEqual(actual, tc.expected) {
		t.Fatalf("TestCase%v failed expected: %v, actual: %v\n", index, tc.expected, actual)
	}
}

func TestCases(t *testing.T) {
	testCases := []TestCase{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 3},
			target:   6,
			expected: []int{0, 2},
		},
	}

	var wg sync.WaitGroup
	for idx, tc := range testCases {
		wg.Add(1)
		go func(idx int, tc TestCase) {
			defer wg.Done()
			tc.Test(idx, t)
		}(idx, tc)
	}
	wg.Wait()
}
