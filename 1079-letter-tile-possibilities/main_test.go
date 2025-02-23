package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	testCase := "AAB"
	expected := 8
	actual := numTilePossibilities(testCase)
	if actual != expected {
		t.Fatalf("TestCase1 failed expected: %v, got: %v", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	testCase := "AAABBC"
	expected := 188
	actual := numTilePossibilities(testCase)
	if actual != expected {
		t.Fatalf("TestCase2 failed expected: %v, got: %v", expected, actual)
	}

}

func TestCase3(t *testing.T) {
	testCase := "V"
	expected := 1
	actual := numTilePossibilities(testCase)
	if actual != expected {
		t.Fatalf("TestCase3 failed expected: %v, got: %v", expected, actual)
	}
}
