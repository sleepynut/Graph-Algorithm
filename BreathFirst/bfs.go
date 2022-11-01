package main

import (
	"fmt"

	u "github.com/sleepynut/Graph-Algorithm/Utils"
)

func bfs(start string, adj map[string][]string) {
	var q u.Queue
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
