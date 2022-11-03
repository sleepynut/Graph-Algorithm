package main

import (
	"fmt"
	"strconv"

	u "github.com/sleepynut/Graph-Algorithm/Utils"
)

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

func traverse(graph map[string][]string, src string, visited *u.Set) {
	if visited.Has(src) {
		return
	}

	// Add current node to visited set
	visited.Add(src)
	for _, neighbor := range graph[src] {
		traverse(graph, neighbor, visited)
	}
}

// count number of connected component recursively
func connectedCompCountRe(graph map[string][]string) int {
	s := u.NewSet()
	count := 0

	for k := range graph {
		if s.Has(k) {
			continue
		}

		traverse(graph, k, &s)
		count++
	}

	return count
}

func traverseCount(graph map[string][]string, src string, visited *u.Set) int {
	if visited.Has(src) {
		return 0
	}

	visited.Add(src)
	count := 1
	for _, neighbor := range graph[src] {
		count += traverseCount(graph, neighbor, visited)
	}

	return count
}

func largestComponent(graph map[string][]string) int {
	largest := 0
	visited := u.NewSet()

	for k := range graph {
		count := traverseCount(graph, k, &visited)

		if count > largest {
			largest = count
		}
	}

	return largest
}

func traverseGrid(grid [][]string, r int, c int, visited *u.Set) {
	row := len(grid)
	col := len(grid[0])

	if r < 0 || r >= row || c < 0 || c >= col {
		return
	}

	pos := strconv.Itoa(r*row + c*col)
	if visited.Has(pos) || grid[r][c] == "W" {
		return
	}

	visited.Add(pos)

	// move up
	traverseGrid(grid, r-1, c, visited)

	// move right
	traverseGrid(grid, r, c+1, visited)

	// move down
	traverseGrid(grid, r+1, c, visited)

	// move left
	traverseGrid(grid, r, c-1, visited)

}

func islandCount(grid [][]string) int {
	count := 0
	row := len(grid)
	col := len(grid[0])
	visited := u.NewSet()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			pos := strconv.Itoa(i*row + j*col)

			if visited.Has(pos) || grid[i][j] == "W" {
				continue
			}

			traverseGrid(grid, i, j, &visited)
			count++
		}
	}

	return count
}

func exploreCount(grid [][]string, r int, c int, visited *u.Set) int {
	row := len(grid)
	col := len(grid[0])

	if r < 0 || r >= row || c < 0 || c >= col {
		return 0
	}

	pos := fmt.Sprintf("%d,%d", r, c)
	if grid[r][c] == "W" || visited.Has(pos) {
		return 0
	}

	visited.Add(pos)
	count := 1

	count += exploreCount(grid, r-1, c, visited)
	count += exploreCount(grid, r, c+1, visited)
	count += exploreCount(grid, r+1, c, visited)
	count += exploreCount(grid, r, c-1, visited)

	return count
}

func minimumIsland(grid [][]string) int {
	row := len(grid)
	col := len(grid[0])
	min := row * col
	visited := u.NewSet()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			count := exploreCount(grid, i, j, &visited)

			if count > 0 && count < min {
				min = count
			}
		}
	}

	return min
}
