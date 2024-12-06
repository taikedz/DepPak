package main

import (
	"fmt"
	"sync"
	"net.taikedz.deppak/deppak/net"
)

func download(number int, url string, wg *sync.WaitGroup, back chan string) {
	defer wg.Done()
	fmt.Println("Starting download ... ", url)
	if err := net.FetchHttp(url, fmt.Sprintf("scratch/file%d.html", number)); err != nil {
		back <- fmt.Sprintf("Failed: %s -> %s", url, err)
		return
	}
	fmt.Println("End download ", url)
}

func main() {
	links := []string{"https://www.example.com", "https://dev.to", "https://www.example.nope"}
	
	// declaring the variable with the struct type initializes a memory space with its zeroed values
    var wg sync.WaitGroup
	wg.Add(len(links))
    messages := make(chan string, len(links))

	for i, url := range links {
		go download(i, url, &wg, messages)
	}

	wg.Wait()
	close(messages)

	for item := range messages {
		fmt.Println(item)
	}
}