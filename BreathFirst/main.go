package main

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
}
