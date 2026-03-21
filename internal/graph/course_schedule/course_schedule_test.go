package graph

import "testing"

func Test_canFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "example 1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "example 2 cycle",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
		{
			name:          "single course no prereq",
			numCourses:    1,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "single course self not in constraints",
			numCourses:    1,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "linear chain",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "triangle cycle",
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {1, 2}, {2, 0}},
			want:          false,
		},
		{
			name:          "diamond no cycle",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			want:          true,
		},
		{
			name:          "empty prerequisites multiple courses",
			numCourses:    5,
			prerequisites: [][]int{},
			want:          true,
		},
		{
			name:          "multiple independent edges",
			numCourses:    3,
			prerequisites: [][]int{{2, 0}, {2, 1}},
			want:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Fatalf("canFinish(%d, %v) = %v; want %v", tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}
