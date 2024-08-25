package main

import (
	"fmt"
	"learn/ch1/gif4"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu = sync.Mutex{}
var count int

func main() {
	http.HandleFunc("/", handler4All)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/gif", fetchGif)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler4All(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "req path: %s", r.URL.Path)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count: %d", count)
	mu.Unlock()
}

func fetchGif(w http.ResponseWriter, r *http.Request) {
	cycles := 5 // default number of cycles
	if err := r.ParseForm(); err == nil {
		if val, ok := r.Form["cycles"]; ok {
			if c, err := strconv.Atoi(val[0]); err == nil {
				cycles = c
			}
		}
	}

	gif.Lissajous(w, float64(cycles))
}
