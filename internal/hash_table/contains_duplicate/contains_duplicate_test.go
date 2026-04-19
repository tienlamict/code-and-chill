package hash_table

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"example 1 - has duplicate", []int{1, 2, 3, 1}, true},
		{"example 2 - all distinct", []int{1, 2, 3, 4}, false},
		{"example 3 - multiple duplicates", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{"single element", []int{1}, false},
		{"two same elements", []int{5, 5}, true},
		{"two different elements", []int{5, 6}, false},
		{"negative numbers duplicate", []int{-1, -2, -3, -1}, true},
		{"negative numbers distinct", []int{-1, -2, -3, -4}, false},
		{"large values", []int{1000000000, -1000000000, 1000000000}, true},
		{"all same", []int{7, 7, 7, 7}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := containsDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
