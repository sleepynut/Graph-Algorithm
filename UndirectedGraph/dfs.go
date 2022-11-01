package main

// dfs with recursive method to check if the given graph has a path from src to dst node
func hasPath(graph map[string][]string, src string, dst string, visited *Set) bool {
	if src == dst {
		return true
	}

	visited.Add(src)
	for _, neighbor := range graph[src] {
		if (*visited)[neighbor] {
			continue
		}

		if hasPath(graph, neighbor, dst, visited) {
			return true
		}
	}

	return false
}
