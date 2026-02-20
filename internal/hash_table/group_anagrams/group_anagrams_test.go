package hash_table

import (
	"reflect"
	"sort"
	"testing"
)

// sortStringSlices sắp xếp các slice trong slice để so sánh
// (vì thứ tự các nhóm và thứ tự trong mỗi nhóm không quan trọng)
func sortStringSlices(slices [][]string) {
	// Sắp xếp từng slice bên trong
	for i := range slices {
		sort.Strings(slices[i])
	}
	// Sắp xếp các slice theo phần tử đầu tiên
	sort.Slice(slices, func(i, j int) bool {
		if len(slices[i]) == 0 && len(slices[j]) == 0 {
			return false
		}
		if len(slices[i]) == 0 {
			return true
		}
		if len(slices[j]) == 0 {
			return false
		}
		return slices[i][0] < slices[j][0]
	})
}

// compareStringSlices so sánh hai slice các slice string
// (không quan tâm đến thứ tự)
func compareStringSlices(got, want [][]string) bool {
	if len(got) != len(want) {
		return false
	}

	// Sắp xếp cả hai để so sánh
	sortStringSlices(got)
	sortStringSlices(want)

	return reflect.DeepEqual(got, want)
}

func Test_GroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "example 1",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name: "example 2 - empty string",
			strs: []string{""},
			want: [][]string{
				{""},
			},
		},
		{
			name: "example 3 - single string",
			strs: []string{"a"},
			want: [][]string{
				{"a"},
			},
		},
		{
			name: "all anagrams",
			strs: []string{"eat", "tea", "ate"},
			want: [][]string{
				{"ate", "eat", "tea"},
			},
		},
		{
			name: "no anagrams",
			strs: []string{"abc", "def", "ghi"},
			want: [][]string{
				{"abc"},
				{"def"},
				{"ghi"},
			},
		},
		{
			name: "multiple groups",
			strs: []string{"abc", "bca", "cab", "xyz", "yzx", "a"},
			want: [][]string{
				{"a"},
				{"abc", "bca", "cab"},
				{"xyz", "yzx"},
			},
		},
		{
			name: "single character strings",
			strs: []string{"a", "b", "a", "b", "c"},
			want: [][]string{
				{"a", "a"},
				{"b", "b"},
				{"c"},
			},
		},
		{
			name: "duplicate strings",
			strs: []string{"eat", "eat", "tea"},
			want: [][]string{
				{"eat", "eat", "tea"},
			},
		},
		{
			name: "long strings",
			strs: []string{"listen", "silent", "enlist"},
			want: [][]string{
				{"enlist", "listen", "silent"},
			},
		},
		{
			name: "empty array",
			strs: []string{},
			want: [][]string{},
		},
		{
			name: "same length different chars",
			strs: []string{"abc", "acb", "bac", "bca", "cab", "cba", "xyz"},
			want: [][]string{
				{"abc", "acb", "bac", "bca", "cab", "cba"},
				{"xyz"},
			},
		},
		{
			name: "different lengths",
			strs: []string{"a", "ab", "ba", "abc", "cba"},
			want: [][]string{
				{"a"},
				{"ab", "ba"},
				{"abc", "cba"},
			},
		},
		{
			name: "complex case",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat", "tab"},
			want: [][]string{
				{"ate", "eat", "tea"},
				{"bat", "tab"},
				{"nat", "tan"},
			},
		},
		{
			name: "repeated characters",
			strs: []string{"aab", "aba", "baa", "abb"},
			want: [][]string{
				{"aab", "aba", "baa"},
				{"abb"},
			},
		},
		{
			name: "very long strings",
			strs: []string{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
			want: [][]string{
				{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Input strs: %v", tt.strs)

			got := groupAnagrams(tt.strs)
			t.Logf("Output: %v", got)

			if !compareStringSlices(got, tt.want) {
				t.Fatalf(
					"groupAnagrams(%v) = %v; want %v",
					tt.strs, got, tt.want,
				)
			}
		})
	}
}
