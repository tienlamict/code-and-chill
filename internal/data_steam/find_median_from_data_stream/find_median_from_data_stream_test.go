package data_stream

import (
	"math"
	"testing"
)

// floatEqual kiểm tra xem hai số float có bằng nhau không (với độ chính xác 10^-5)
func floatEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-5
}

func Test_MedianFinder(t *testing.T) {

	tests := []struct {
		name     string
		operations []string
		values   []interface{}
		want     []interface{}
	}{
		{
			name:     "example 1",
			operations: []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			values:   []interface{}{nil, 1, 2, nil, 3, nil},
			want:     []interface{}{nil, nil, nil, 1.5, nil, 2.0},
		},
		{
			name:     "single number",
			operations: []string{"MedianFinder", "addNum", "findMedian"},
			values:   []interface{}{nil, 5, nil},
			want:     []interface{}{nil, nil, 5.0},
		},
		{
			name:     "two numbers",
			operations: []string{"MedianFinder", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, 1, 2, nil},
			want:     []interface{}{nil, nil, nil, 1.5},
		},
		{
			name:     "three numbers",
			operations: []string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, 1, 2, 3, nil},
			want:     []interface{}{nil, nil, nil, nil, 2.0},
		},
		{
			name:     "increasing sequence",
			operations: []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, 1, 2, 3, 4, 5, nil},
			want:     []interface{}{nil, nil, nil, nil, nil, nil, 3.0},
		},
		{
			name:     "decreasing sequence",
			operations: []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, 5, 4, 3, 2, 1, nil},
			want:     []interface{}{nil, nil, nil, nil, nil, nil, 3.0},
		},
		{
			name:     "mixed sequence",
			operations: []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, 3, 1, 4, 2, nil},
			want:     []interface{}{nil, nil, nil, nil, nil, 2.5},
		},
		{
			name:     "negative numbers",
			operations: []string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, -1, -2, -3, nil},
			want:     []interface{}{nil, nil, nil, nil, -2.0},
		},
		{
			name:     "duplicate numbers",
			operations: []string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, 2, 2, 2, nil},
			want:     []interface{}{nil, nil, nil, nil, 2.0},
		},
		{
			name:     "large numbers",
			operations: []string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, 100000, 200000, 300000, nil},
			want:     []interface{}{nil, nil, nil, nil, 200000.0},
		},
		{
			name:     "multiple findMedian calls",
			operations: []string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian", "addNum", "findMedian"},
			values:   []interface{}{nil, 1, nil, 2, nil, 3, nil},
			want:     []interface{}{nil, nil, 1.0, nil, 1.5, nil, 2.0},
		},
		{
			name:     "even count median",
			operations: []string{"MedianFinder", "addNum", "addNum", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, 1, 2, 3, 4, nil},
			want:     []interface{}{nil, nil, nil, nil, nil, 2.5},
		},
		{
			name:     "odd count median",
			operations: []string{"MedianFinder", "addNum", "addNum", "addNum", "findMedian"},
			values:   []interface{}{nil, 1, 2, 3, nil},
			want:     []interface{}{nil, nil, nil, nil, 2.0},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			var mf *MedianFinder

			for i, op := range tt.operations {

				switch op {
				case "MedianFinder":
					m := Constructor()
					mf = &m
					if tt.want[i] != nil {
						t.Fatalf("Expected nil for MedianFinder constructor")
					}

				case "addNum":
					if mf == nil {
						t.Fatalf("MedianFinder not initialized")
					}
					num := tt.values[i].(int)
					mf.AddNum(num)
					if tt.want[i] != nil {
						t.Fatalf("Expected nil for addNum")
					}

				case "findMedian":
					if mf == nil {
						t.Fatalf("MedianFinder not initialized")
					}
					got := mf.FindMedian()
					want := tt.want[i].(float64)

					t.Logf("After operations up to index %d, findMedian() = %v", i, got)

					if !floatEqual(got, want) {
						t.Fatalf(
							"findMedian() = %v; want %v",
							got, want,
						)
					}
				}
			}

		})

	}

}
