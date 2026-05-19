package sum_of_two_integers

import "testing"

func TestGetSum(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Example 1", 1, 2, 3},
		{"Example 2", 2, 3, 5},
		{"Zero + Positive", 0, 5, 5},
		{"Positive + Zero", 10, 0, 10},
		{"Two Zeros", 0, 0, 0},
		{"Negative + Positive", -1, 1, 0},
		{"Negative + Negative", -2, -3, -5},
		{"Positive + Large Negative", 10, -15, -5},
		{"Boundary Case -1000 + 1000", -1000, 1000, 0},
		{"Maximum constraints", 1000, 1000, 2000},
		{"Minimum constraints", -1000, -1000, -2000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getSum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("getSum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
