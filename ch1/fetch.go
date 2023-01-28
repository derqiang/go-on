package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// Fetch : Page 38, Chapter 1
func Fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf(" On Error : %v \n", err)
			os.Exit(1)
		}

		// Example
		//body, err := ioutil.ReadAll(resp.Body)
		//_ = resp.Body.Close()
		//if err != nil {
		//	fmt.Printf(" Read body from http response occured error , %v \n", err)
		//	os.Exit(1)
		//}
		//fmt.Printf("Success and the response is \n%v \n", string(body))

		// Practice 1.7
		r, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("io.Copy call error , %v \n", err)
			os.Exit(1)
		}
		fmt.Printf("\nio.Copy mehtod output count : %v \n", r)
		_ = resp.Body.Close()
	}
}

// FetchAll , Page 40, Chapter 1.6 并发获取多个URL
// 2022-07-19 记录： 可以批量多次请求同一个网站，以返现其缓存的支持与否
func FetchAll() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go tFetch(url, ch)
		fmt.Println(url)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func tFetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	defer func() {
		fmt.Println(" 释放资源 ...")
		resp.Body.Close()
	}()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
