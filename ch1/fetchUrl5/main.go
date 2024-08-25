package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	urls := os.Args[1:]

	for _, url := range urls {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		//body, err := json.Marshal(resp.Body)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", url, err)
			//fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(resp.StatusCode)
		fmt.Println(string(body))
	}

}
