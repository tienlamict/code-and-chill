package hash_table

import (
	"reflect"
	"testing"
)

// buildGraph creates a graph from adjacency list representation
// adjList[i] contains neighbors of node (i+1)
func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList))
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1, Neighbors: make([]*Node, 0)}
	}

	for i, neighbors := range adjList {
		for _, neighborVal := range neighbors {
			// neighborVal is 1-indexed, convert to 0-indexed
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighborVal-1])
		}
	}

	return nodes[0]
}

// graphToAdjList converts graph to adjacency list representation
func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}

	visited := make(map[*Node]bool)
	adjList := make([][]int, 0)

	var dfs func(*Node)
	dfs = func(n *Node) {
		if visited[n] {
			return
		}
		visited[n] = true

		// Extend adjList if needed (nodes are 1-indexed)
		for len(adjList) < n.Val {
			adjList = append(adjList, []int{})
		}

		// Collect neighbor values
		neighbors := make([]int, 0, len(n.Neighbors))
		for _, neighbor := range n.Neighbors {
			neighbors = append(neighbors, neighbor.Val)
			dfs(neighbor)
		}

		adjList[n.Val-1] = neighbors
	}

	dfs(node)
	return adjList
}

// isSameGraphStructure checks if two graphs have the same structure
// by comparing their adjacency lists
func isSameGraphStructure(node1 *Node, node2 *Node) bool {
	adjList1 := graphToAdjList(node1)
	adjList2 := graphToAdjList(node2)
	return reflect.DeepEqual(adjList1, adjList2)
}

// areGraphsDifferent checks if two graphs are different objects (different memory addresses)
func areGraphsDifferent(node1 *Node, node2 *Node) bool {
	visited1 := make(map[*Node]bool)
	visited2 := make(map[*Node]bool)

	var dfs func(*Node, *Node) bool
	dfs = func(n1 *Node, n2 *Node) bool {
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}

		// Check if nodes are different objects (different memory addresses)
		if n1 == n2 {
			return false
		}

		if visited1[n1] || visited2[n2] {
			return true
		}

		visited1[n1] = true
		visited2[n2] = true

		if n1.Val != n2.Val {
			return false
		}

		if len(n1.Neighbors) != len(n2.Neighbors) {
			return false
		}

		for i := range n1.Neighbors {
			if !dfs(n1.Neighbors[i], n2.Neighbors[i]) {
				return false
			}
		}

		return true
	}

	return dfs(node1, node2)
}

func Test_CloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
		want    [][]int
	}{
		{
			name:    "example 1",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			want:    [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "example 2 - single node",
			adjList: [][]int{{}},
			want:    [][]int{{}},
		},
		{
			name:    "example 3 - empty graph",
			adjList: [][]int{},
			want:    [][]int{},
		},
		{
			name:    "two connected nodes",
			adjList: [][]int{{2}, {1}},
			want:    [][]int{{2}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			t.Logf("Input adjList: %v", tt.adjList)

			cloned := cloneGraph(original)

			if original == nil {
				if cloned != nil {
					t.Fatalf("cloneGraph(nil) = %v; want nil", cloned)
				}
				return
			}

			// Convert cloned graph back to adjacency list
			got := graphToAdjList(cloned)
			t.Logf("Output adjList: %v", got)

			// Verify structure is the same
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("cloneGraph(%v) produced adjList %v; want %v", tt.adjList, got, tt.want)
			}

			// Verify graphs are different objects (deep copy)
			if !areGraphsDifferent(original, cloned) {
				t.Fatalf("cloneGraph did not create a deep copy - original and cloned share memory")
			}

			// Verify structure matches
			if !isSameGraphStructure(original, cloned) {
				t.Fatalf("Cloned graph structure does not match original")
			}
		})
	}
}
