package main

import "fmt"

func main() {
	s := NewSet()

	s.Add("a")
	s.Add("a")
	s.Add("b")
	s.Remove("x")
	s.Remove("a")
	s.Display()

	fmt.Println()

	edges := [][]string{
		{"i", "j"},
		{"k", "i"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}

	graph := buildGraph(edges)

	for k, v := range graph {
		fmt.Printf("%s: %v\n", k, v)
	}

	src := "j"
	dst := "m"
	visited := NewSet()

	fmt.Printf("hasPath from %s - %s: %t\n", src, dst, hasPath(graph, src, dst, &visited))

}

func buildGraph(edges [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	return graph
}
