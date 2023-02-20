package main

import (
	"fmt"
	"sort"
)

type Prereqs map[string][]string

var prereqs = Prereqs{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func (p Prereqs) Run() {
	for i, item := range topoSort(prereqs) {
		fmt.Printf("%d: \t%s\n", i+1, item)
	}
}

func topoSort(m Prereqs) []string {
	var orders []string
	var seen = make(map[string]int)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			seen[item]++
			if seen[item] == 1 {
				visitAll(m[item])
				orders = append(orders, item)
			} else if seen[item] == 2 {
				fmt.Printf("cycle item : %v\n", item)
			}
		}
	}
	var keys []string
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	fmt.Printf("seen items : %v\n", seen)
	return orders
}
