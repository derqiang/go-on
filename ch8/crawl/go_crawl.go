package main

import (
	"fmt"
	"go_try/me"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := me.Extract(url)
	if err != nil {
		log.Println(err)
	}
	return list
}

var tokens = make(chan struct{}, 20)

func crawlLimit(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := me.Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}
	return list
}

func main() {
	//limitRunning1()
	limitRunning2()
}

func limitRunning1() {
	worklist := make(chan []string)
	var n int

	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	var seens = make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seens[link] {
				seens[link] = true
				n++
				go func(link string) {
					worklist <- crawlLimit(link)
				}(link)
			}
		}
	}
}

type Link struct {
	link  string
	depth int
}

func packLinks(urls []string, depth int) (links []Link) {
	for _, in := range urls {
		links = append(links, Link{in, depth})
	}
	return links
}

func limitRunning2() {

	worklist := make(chan []Link)
	unseenLinks := make(chan Link)
	go func() { worklist <- packLinks(os.Args[1:], 0) }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := packLinks(crawl(link.link), link.depth+1)
				go func() { worklist <- foundLinks }()
				if link.depth > 2 {
					log.Println("\t\t深度已达到，退出爬取任务")
					return
				}
			}
		}()
	}

	var seens = make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seens[link.link] {
				seens[link.link] = true
				unseenLinks <- link
			}
		}
	}
}
