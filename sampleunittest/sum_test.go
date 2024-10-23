package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("2 and 3, should return 5", func(t *testing.T) {
		result := Sum(2, 3)
		assert.Equal(t, 5, result)
	})

	t.Run("3 and 5, should return 8", func(t *testing.T) {
		result := Sum(3, 5)
		assert.Equal(t, 8, result)
	})
}

func TestSum2(t *testing.T) {
	testCase := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "2 and 3, should return 5",
			a:        2,
			b:        3,
			expected: 5,
		},
		{
			name:     "3 and 5, should return 8",
			a:        3,
			b:        5,
			expected: 8,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result := Sum(tc.a, tc.b)
			assert.Equal(t, tc.expected, result)
		})
	}
}
