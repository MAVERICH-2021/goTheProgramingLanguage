package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

func main() {
	//absPath, _ := os.Getwd()
	absPath := strings.Replace("C:\\Users\\maver\\GolandProjects\\theProgrammingLanguage\\ch1",
		"\\", "/", -1)
	counter := make(map[string]int) // make create a reference
	//counter := map[string]int{"sss": 1}
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		fmt.Println("Please give me the file names")
	} else {
		for _, fileName := range fileNames {
			filePath := path.Join(absPath, fileName)
			file, err := os.Open(filePath)
			if err != nil {
				print(err.Error())
				continue
			}
			countLines(file, counter)
			file.Close()
		}
	}

	for k, v := range counter {
		if v > 0 {
			fmt.Println(k)
		}
	}

}

func countLines(file *os.File, counter map[string]int) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ling := scanner.Text()
		counter[ling]++
	}
	print(json.Marshal(counter))
}
