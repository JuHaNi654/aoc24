package main

import (
	"fmt"
	"testing"
)

func TestIsIncreasing(t *testing.T) {
	cases := []struct {
		Input    []int
		Expected bool
	}{
		{Input: []int{8, 1, 2, 3, 4}, Expected: true},
		{Input: []int{1, 2, 3, 4, 2}, Expected: true},
		{Input: []int{1, 3, 2, 4, 5}, Expected: true},
		{Input: []int{1, 3, 3, 4, 5}, Expected: true},
		{Input: []int{1, 1, 2, 3, 4}, Expected: true},
		{Input: []int{1, 2, 3, 4, 4}, Expected: true},
		{Input: []int{27, 28, 30, 30, 32, 35}, Expected: true},

		{Input: []int{48, 46, 47, 49, 51, 54, 56}, Expected: true},
		{Input: []int{1, 1, 2, 3, 4, 5}, Expected: true},
		{Input: []int{1, 2, 3, 4, 5, 5}, Expected: true},
		{Input: []int{5, 1, 2, 3, 4, 5}, Expected: true},
		{Input: []int{1, 6, 7, 8, 9, 10}, Expected: true},

		{Input: []int{1, 2, 3, 4, 5}, Expected: true},
		{Input: []int{8, 2, 3, 4, 5}, Expected: true},
		{Input: []int{1, 8, 3, 4, 5}, Expected: true},
		{Input: []int{1, 2, 8, 4, 5}, Expected: true},
		{Input: []int{1, 2, 3, 8, 5}, Expected: true},
		{Input: []int{1, 2, 3, 4, 8}, Expected: true},

		{Input: []int{81, 83, 85, 881, 92}, Expected: false},

		{Input: []int{1, 1, 1, 1, 1}, Expected: false},
		{Input: []int{1, 8, 2, 3, 8}, Expected: false},
		{Input: []int{8, 1, 2, 3, 8}, Expected: false},
		{Input: []int{1, 3, 2, 3, 1}, Expected: false},
		{Input: []int{1, 2, 3, 2, 1}, Expected: false},
		{Input: []int{1, 2, 3, 8, 9}, Expected: false},
		{Input: []int{1, 1, 2, 2, 3}, Expected: false},
		{Input: []int{80, 78, 83, 84, 87, 89, 93}, Expected: false},
		{Input: []int{15, 12, 15, 18, 20, 23, 25, 27}, Expected: true},
		{Input: []int{90, 89, 91, 93, 95, 94}, Expected: false},
		{Input: []int{9, 12, 9, 10}, Expected: false},
		{Input: []int{9, 12, 10, 12}, Expected: true},
		{Input: []int{9, 12, 5}, Expected: true},
		{Input: []int{15, 12, 15, 18, 20, 23, 25, 27}, Expected: true},
		{Input: []int{10, 13, 12, 13, 14}, Expected: true},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%v", test.Input), func(t *testing.T) {
			result := isSafe(test.Input, 0, 1, true, Asc)
			if result != test.Expected {
				t.Errorf("got %v, want %v", result, test.Expected)
			}
		})
	}
}

func TestIsDecreasing(t *testing.T) {
	cases := []struct {
		Input    []int
		Expected bool
	}{

		{Input: []int{58, 55, 53, 52, 51, 49, 46, 42}, Expected: true},
		{Input: []int{58, 55, 53, 52, 51, 49, 70, 46}, Expected: true},
		{Input: []int{63, 58, 55, 53, 52, 51, 49, 46}, Expected: true},
		{Input: []int{58, 70, 55, 53, 52, 51, 49, 46}, Expected: true},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%v", test.Input), func(t *testing.T) {
			result := isSafe(test.Input, 0, 1, true, Desc)
			if result != test.Expected {
				t.Errorf("got %v, want %v", result, test.Expected)
			}
		})
	}
}
