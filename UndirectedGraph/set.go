package main

import "fmt"

type Set map[string]bool

func (s *Set) Has(node string) bool {
	return (*s)[node]
}

func (s *Set) Add(node string) {
	(*s)[node] = true
}

func (s *Set) Remove(node string) {
	if s.Has(node) {
		delete(*s, node)
	}
}

func (s *Set) Display() {
	for k := range *s {
		fmt.Printf("%s, ", k)
	}
}

func NewSet() Set {
	return make(map[string]bool)
}
