package main

import (
	"fmt"

	u "github.com/sleepynut/Graph-Algorithm/Utils"
)

func main() {
	adj := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	bfs("a", adj)
	src := "d"
	dst := "d"
	fmt.Printf("Has Path from %s - %s: %t", src, dst, hasPath(adj, src, dst))
}

func hasPath(graph map[string][]string, src string, dst string) bool {
	var (
		q  u.Queue
		el string
	)

	q.Push(src)

	for !q.IsEmpty() {
		el, _ = q.Pop()
		if el == dst {
			return true
		}

		for _, neighbor := range graph[el] {
			q.Push(neighbor)
		}
	}

	return false
}
