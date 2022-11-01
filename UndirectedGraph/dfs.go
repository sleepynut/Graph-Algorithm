package main

import u "github.com/sleepynut/Graph-Algorithm/Utils"

// dfs with recursive method to check if the given graph has a path from src to dst node
func hasPath(graph map[string][]string, src string, dst string, visited *u.Set) bool {
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

func connectedCompCount(graph map[string][]string) int {
	s := u.NewSet()
	var stack u.Stack
	var count int

	//traverse the graph using depth first method
	for k := range graph {

		// if the current node has already been traversed
		// skip this traversal using current node as source
		if s.Has(k) {
			continue
		}

		stack.Push(k)
		// traverse graph
		for !stack.IsEmpty() {
			src, _ := stack.Pop()

			// check if current node has been visited
			// yes - skip
			// no - continue this traversal
			if s.Has(src) {
				continue
			}

			// mark current node as visited
			s.Add(src)

			// Enlisting all neighbor of current node
			for _, neighbor := range graph[src] {
				stack.Push(neighbor)
			}
		}

		count++
	}

	return count
}
