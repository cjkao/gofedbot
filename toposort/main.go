package main

import "fmt"

var preq = map[string][]string{
	"algo":           {"a", "data structure"},
	"calculaus":      {"linear algebra", "d"},
	"compilers":      {"data structure", "formal languages", "computer organization"},
	"data structure": {"discrete math"},
	"discrete math":  {"d"},
	"network":        {"os", "algo"},
	"os":             {"computer organization", "data structure"},
}

func topoSort(m map[string][]string) []string {
	var visitAll func([]string)
	seen := make(map[string]bool)
	var order []string
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}

	}

	visitAll(getKeys(m))
	return order
}
func getKeys(m map[string][]string) []string {
	keys := make([]string, len(m))
	j := 0
	for key := range m {
		keys[j] = key
		j++
	}
	return keys
}
func main() {
	for i, course := range topoSort(preq) {
		fmt.Printf("%d %s\n", i, course)
	}
}
