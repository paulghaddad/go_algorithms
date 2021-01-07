package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsValidSubsequence(t *testing.T) {
	var tests = []struct {
		array    []int
		sequence []int
		want     bool
	}{
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10, 12}, false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v, %v", tt.array, tt.sequence)
		t.Run(testName, func(t *testing.T) {
			got := IsValidSubsequence(tt.array, tt.sequence)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
