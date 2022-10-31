package main

import "fmt"

func bfd(start string, adj map[string][]string) {
	var q Queue
	q.Push(start)

	for {
		v, b := q.Pop()
		if !b {
			break
		}

		// print current node
		fmt.Print(v + ", ")

		for _, node := range adj[v] {
			q.Push(node)
		}
	}

	fmt.Println()
}
