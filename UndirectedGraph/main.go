package main

import (
	"fmt"

	u "github.com/sleepynut/Graph-Algorithm/Utils"
)

func main() {
	s := u.NewSet()

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
	visited := u.NewSet()

	fmt.Printf("hasPath from %s - %s: %t\n", src, dst, hasPath(graph, src, dst, &visited))

	graph2 := map[string][]string{
		"0": {"8", "1", "5"},
		"1": {"0"},
		"5": {"0", "8"},
		"8": {"0", "5"},
		"2": {"3", "4"},
		"3": {"2", "4"},
		"4": {"3", "2"},
	}

	fmt.Printf("Connected Component Count: %d\n", connectedCompCount(graph2))
	fmt.Printf("Connected Component Count (Re): %d\n", connectedCompCountRe(graph2))

	fmt.Printf("Largest component count: %d\n", largestComponent(graph2))

	edges2 := [][]string{
		{"w", "x"},
		{"x", "y"},
		{"z", "y"},
		{"z", "v"},
		{"w", "v"},
	}

	graph3 := buildGraph(edges2)
	src = "w"
	dst = "z"
	fmt.Printf("Shortest path from %s - %s: %d\n", src, dst, shortestPath(graph3, src, dst))
}

func buildGraph(edges [][]string) map[string][]string {
	graph := make(map[string][]string)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	return graph
}
