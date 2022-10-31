package main

import "fmt"

func dfs(start string, adj map[string][]string) {
	var s Stack
	s.Push(start)

	for true {

		v, b := s.Pop()
		if !b {
			break
		}

		// print current node
		fmt.Println(v + ", ")

		for _, node := range adj[v] {
			s.Push(node)
		}

	}
}
