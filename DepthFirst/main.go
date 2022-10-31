package main

func main() {
	adj := map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	dfs("a", adj)
}
