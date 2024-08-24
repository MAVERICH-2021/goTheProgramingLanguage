package main

import (
	"fmt"
	"os"
	"strings"
)

// fetch cli args
func main() {
	//var s, sep string
	// normal loop
	//for i := 0; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}

	// range
	//for _, i := range os.Args[1:] { // key value pair
	//	s += sep + i
	//	sep = " "
	//}
	//fmt.Println(s)
	fmt.Println(os.Args[0]) //-> abs path of this file
	fmt.Println(strings.Join(os.Args[1:], " "))
}
