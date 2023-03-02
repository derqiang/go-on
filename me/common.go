package me

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

type Runner interface {
	Run()
}

func ForEachNode(doc *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(doc)
	}
	for n := doc.FirstChild; n != nil; n = n.NextSibling {
		ForEachNode(n, pre, post)
	}
	if post != nil {
		post(doc)
	}
}

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	var visitNode func(node *html.Node)
	visitNode = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	ForEachNode(doc, visitNode, nil)
	return links, nil
}

func BreathFirst(f func(item string) []string, workList []string) {
	seen := make(map[string]bool)
	for len(workList) > 0 {
		items := workList
		workList = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				workList = append(workList, f(item)...)
			}
		}
	}
}
