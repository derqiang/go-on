package main

import (
	"bufio"
	"fmt"
	"github.com/dlclark/regexp2"
	"os"
)

func main() {
	//abp, _ := filepath.Abs("./")
	//filepath.Walk("./", func(p string, info fs.FileInfo, err error) error {
	//	if info.IsDir() && p != "./" {
	//		return filepath.SkipDir
	//	}
	//	if strings.HasSuffix(info.Mode().String(), "x") {
	//		//fmt.Printf(" > %v \n ", path.Join("./", info.Name()))
	//		if err := os.Remove(info.Name()); err != nil {
	//			fmt.Fprintf(os.Stderr, " errors : %v\n", err)
	//		}
	//	}
	//	return nil
	//})
	//return
	for _, p := range os.Args[1:] {
		f, err := os.OpenFile(p, os.O_RDWR, 0777)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read file error : %v\n", err)
		}
		s := bufio.NewScanner(f)
		var n int
		for s.Scan() {
			n++
			reg := regexp2.MustCompile(`(?<=listen.+)(\d{4})`, 0)
			content, err := reg.FindStringMatch(s.Text())
			if err != nil {
				fmt.Fprintf(os.Stderr, "error happened: %v\n", err)
			}
			if content != nil {
				fmt.Printf(" %d >>> %v - %v\n", n, s.Text(), content.GroupByNumber(1).String())
			}
		}
	}
}

func modifyFile() {
	for _, p := range os.Args[1:] {
		f, err := os.OpenFile(p, os.O_RDWR, 0777)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read file error : %v\n", err)
		}
		//content := f.Read()
	}
}
