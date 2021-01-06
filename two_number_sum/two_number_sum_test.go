package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoNumberSum(t *testing.T) {
	var tests = []struct {
		array     []int
		targetSum int
		want      []int
	}{
		{[]int{3, 5, -4, 8, 11, 1, -1, 6}, 10, []int{-1, 11}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.array, tt.targetSum)
		t.Run(testName, func(t *testing.T) {
			got := TwoNumberSum(tt.array, tt.targetSum)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v; want: %v", got, tt.want)
			}
		})
	}
}
