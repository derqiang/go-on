package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

type Runner interface {
	Run()
}

type FindLinkRunner int64
type FindLinkRunner2 int64
type FindLinkRunner3 int64

func (flr FindLinkRunner) Run() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func (runner FindLinkRunner2) Run() {
	for _, url := range os.Args[1:] {
		links, err := findLink2(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "FindLinkRunner2 Get Url Error : %v\n", err)
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func (runner FindLinkRunner3) Run() {
	BreathFirst(Crawl, os.Args[1:])
}

// 直接用被调用的函数多值返回
func findLinksLog(url string) ([]string, error) {
	log.Printf("findLinks %s", url)
	return findLink2(url)
}

func findLink2(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit(links, c)
	//}
	return traverse(n.FirstChild, links)
}

func traverse(c *html.Node, links []string) []string {
	if c == nil {
		return links
	}
	links = visit(links, c)
	return traverse(c.NextSibling, links)
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

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
