package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var counter = map[string]int{}

type OutlineRunner int64

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		counter[n.Data]++
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func outline2(n *html.Node, pre, post func(node *html.Node)) {
	if n != nil {
		pre(n)
	}
	for ch := n.FirstChild; ch != nil; ch = ch.NextSibling {
		outline2(ch, pre, post)
	}
	if n != nil {
		post(n)
	}
}

func (runner OutlineRunner) Run() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	//outline(nil, doc)
	//fmt.Printf(">>>%*s", 20, "VVV")
	var depth int
	outline2(doc, func(node *html.Node) {
		if node.Type == html.ElementNode {
			if node.FirstChild == nil {
				fmt.Printf("%*s<%s/>\n", depth*2, "", node.Data)
			} else {
				fmt.Printf("%*s<%s>\n", depth*2, "", node.Data)
			}
			depth++
		}
	}, func(node *html.Node) {
		if node.Type == html.ElementNode {
			depth--
			if node.FirstChild != nil {
				fmt.Printf("%*s</%s>\n", depth*2, "", node.Data)
			}
		}
	})
	fmt.Printf("Counter : %v\n", counter)
}
