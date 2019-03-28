package main

import (
	"os"
	"fmt"
)

func main() {
	for k,v := range os.Args {
		fmt.Println("k=", k, " v=", v)
	}
}
