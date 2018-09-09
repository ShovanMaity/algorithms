package samplepackage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jumpSearch(t *testing.T) {
	tests := []struct {
		name          string
		array         []int64
		key           int64
		expectedValue int
	}{
		{
			name:          "When key is present in array",
			array:         []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			key:           1,
			expectedValue: 0,
		},
		{
			name:          "When key is not present in array",
			array:         []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			key:           10,
			expectedValue: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualValue := jumpSearch(tt.array, tt.key)
			assert.Equal(t, tt.expectedValue, actualValue)
		})
	}
}
