package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

// CountWordsAndImages 如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。(通过返回值的命名识别要返回的变量)这称之为bare return。（P174）
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML : %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (word, image int) {
	return 10, 11
}
