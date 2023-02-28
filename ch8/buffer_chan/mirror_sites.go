package main

func main() {
	responses := make(chan string, 3)
	go func() { responses <- request("asia.gopl.io") }()
	go func() { responses <- request("europe.gopl.io") }()
	go func() { responses <- request("americas.gopl.io") }()

	// 卡住慢的，只打印出最快的
	println(<-responses)
}

func request(hostname string) (res string) {
	return "1"
}
