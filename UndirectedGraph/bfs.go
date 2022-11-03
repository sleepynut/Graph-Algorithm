package main

import u "github.com/sleepynut/Graph-Algorithm/Utils"

func shortestPath(graph map[string][]string, src string, dst string) int {
	var qT QueueT[*Passed]
	visited := u.NewSet()
	dist := -1

	qT.Push(&Passed{Node: src, Dist: 0})

	for !qT.IsEmpty() {
		current, _ := qT.Pop()

		if visited[current.Node] {
			continue
		}

		visited.Add(current.Node)

		if current.Node == dst {
			dist = current.Dist
			break
		}

		for _, neighbor := range graph[current.Node] {
			qT.Push(&Passed{Node: neighbor, Dist: current.Dist + 1})
		}
	}

	return dist
}
