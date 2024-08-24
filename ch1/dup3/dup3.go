package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counter := map[string]int{}
	fileNames := os.Args[1:]
	for _, fileName := range fileNames {
		fBytes, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		content := string(fBytes)
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			counter[line]++
		}
	}

	for k, v := range counter {
		//fmt.Println(k, ":", v)
		if v > 0 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}

}
