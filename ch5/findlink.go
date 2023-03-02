package main

import (
	"fmt"
	"go_try/me"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

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
	me.BreathFirst(Crawl, os.Args[1:])
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

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := me.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
