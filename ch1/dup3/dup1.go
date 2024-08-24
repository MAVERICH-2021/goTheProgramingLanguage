package main

import (
	"bufio"
	"fmt"
	"os"
)

// basic io reader
func main() {
	counter := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {
			break
		}
		counter[scanner.Text()]++
	}
	for i, n := range counter {
		if n > 1 {
			fmt.Println(i, ":", n)
		}
	}
}
