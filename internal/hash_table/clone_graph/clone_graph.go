package hash_table

// Node represents a node in an undirected graph
type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph returns a deep copy of the graph using DFS
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	// Map to store original node -> cloned node mapping
	visited := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(original *Node) *Node {
		// If node already cloned, return the clone
		if cloned, exists := visited[original]; exists {
			return cloned
		}

		// Create a new node with the same value
		cloned := &Node{
			Val:       original.Val,
			Neighbors: make([]*Node, 0),
		}

		// Mark this node as visited/cloned
		visited[original] = cloned

		// Recursively clone all neighbors
		for _, neighbor := range original.Neighbors {
			cloned.Neighbors = append(cloned.Neighbors, dfs(neighbor))
		}

		return cloned
	}

	return dfs(node)
}
