package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

//todo understand file system in go
//todo quest: why go cannot open a relative path file

func main() {
	filePath := strings.Replace("C:\\Users\\maver\\GolandProjects\\theProgrammingLanguage\\ch1\\dup1.txt",
		"\\", "/", -1)
	//filename := "dup1.txt"
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//print file content
	fmt.Println("Dup1 file opened")
}
