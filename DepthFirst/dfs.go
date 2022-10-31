package main

import "fmt"

func dfs(start string, adj map[string][]string) {
	var s Stack
	s.Push(start)

	for {

		v, b := s.Pop()
		if !b {
			break
		}

		// print current node
		fmt.Print(v + ", ")

		for _, node := range adj[v] {
			s.Push(node)
		}

	}
	fmt.Println()
}

func dfsRe(current string, adj map[string][]string) {
	// print current node
	fmt.Print(current + ", ")

	if len(adj[current]) == 0 {
		return
	}

	for _, node := range adj[current] {
		dfsRe(node, adj)
	}
}
