package main

import "fmt"

func main() {
	adj := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	dfs("a", adj)
	dfsRe("a", adj)

	src := "a"
	dst := "a"
	fmt.Printf("\nHas Path from %s - %s: %t\n", src, dst, hasPath(adj, src, dst))
	fmt.Printf("\nHas Path (Re) from %s - %s: %t\n", src, dst, hasPathRe(adj, src, dst))
}

func hasPath(graph map[string][]string, src string, dst string) bool {
	var s Stack
	var v string
	var b bool
	s.Push(src)

	for {
		v, b = s.Pop()
		if !b || v == dst {
			break
		}

		for _, node := range graph[v] {
			s.Push(node)
		}
	}

	return v == dst
}

func hasPathRe(graph map[string][]string, src string, dst string) bool {
	if src == dst {
		return true
	}

	for _, neighbor := range graph[src] {
		if hasPathRe(graph, neighbor, dst) {
			return true
		}
	}

	return false
}
