package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	urls := os.Args[1:]
	ch := make(chan string)
	for _, url := range urls {
		fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	//start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- err.Error()
		return
	}
	defer resp.Body.Close()
	nbytes, err := io.Copy(io.Discard, resp.Body) // return byte count - discard content
	if err != nil {
		ch <- err.Error()
		return
	}

	//todo why this shit not work
	//secs := time.Since(start).Seconds()
	ch <- strconv.FormatInt(nbytes, 10) + "  " //fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
