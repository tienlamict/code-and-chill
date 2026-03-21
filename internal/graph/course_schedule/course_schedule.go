package graph

// canFinish trả về true nếu có thể hoàn thành hết numCourses khóa học.
//
// prerequisites[i] = [ai, bi] nghĩa là phải học bi trước mới được học ai
// → cạnh có hướng bi -> ai (topological order phải đi qua bi trước ai).
//
// Bài toán tương đương: đồ thị có hướng có chu trình hay không?
// - Nếu có chu trình → không thể xếp thứ tự học → false.
// - Nếu không (DAG) → có topological sort → true.
//
// Thuật toán: Kahn (BFS) — đếm bậc vào (indegree), liên tục lấy các đỉnh indegree = 0,
// bỏ cạnh đi ra, nếu cuối cùng đã "thăm" đủ numCourses đỉnh thì không có chu trình.
//
// Độ phức tạp: O(V + E) với V = numCourses, E = len(prerequisites).
func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, p := range prerequisites {
		a, b := p[0], p[1] // học b trước a
		adj[b] = append(adj[b], a)
		indegree[a]++
	}

	q := make([]int, 0, numCourses)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	taken := 0
	for head := 0; head < len(q); head++ {
		u := q[head]
		taken++
		for _, v := range adj[u] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return taken == numCourses
}
