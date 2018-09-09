package samplepackage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bubbleSort(t *testing.T) {
	tests := []struct {
		name          string
		array         []int64
		expectedArray []int64
	}{
		{
			name:          "Elements are in random order",
			array:         []int64{1, 8, 3, 4, 6, 5, 7, 2, 9},
			expectedArray: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:          "Elements are in sorted order",
			array:         []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedArray: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:          "Elements are in reverse sorted order",
			array:         []int64{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expectedArray: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValue := bubbleSort(tt.array)
			assert.Equal(t, tt.expectedArray, actualValue)
		})
	}
}

func Test_recursiveBubbleSort(t *testing.T) {
	tests := []struct {
		name          string
		array         []int64
		expectedArray []int64
	}{
		{
			name:          "Elements are in random order",
			array:         []int64{1, 8, 3, 4, 6, 5, 7, 2, 9},
			expectedArray: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:          "Elements are in sorted order",
			array:         []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedArray: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:          "Elements are in reverse sorted order",
			array:         []int64{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expectedArray: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recursiveBubbleSort(tt.array, len(tt.array))
			assert.Equal(t, tt.expectedArray, tt.array)
		})
	}
}
