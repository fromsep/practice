package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) < 1 {
		fmt.Println("Please select 1 file at lest!")
		os.Exit(1)
	}

	for _,file := range files {
		fp, err := os.Open(file)
		if err != nil {
			fmt.Println("The file '" + file + "' open fail!")
			os.Exit(1)
		}
		counter := countLines(fp)
		if check(counter) == false {
			fmt.Println("The file '" + file + "' has duplicate lines!")
		}
	}
}

func countLines(fp *os.File) map[string]int {
	lineCounter := make(map[string]int)
	input := bufio.NewScanner(fp)
	for input.Scan() {
		lineCounter[input.Text()]++
	}
	return lineCounter
}

func check(counter map[string]int) bool  {
	if len(counter) <= 0 {
		return true
	}

	for _, num := range counter {
		if num > 1 {
			return false
		}
	}

	return true
}
