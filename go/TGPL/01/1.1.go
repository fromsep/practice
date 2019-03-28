package main

import(
	"os"
	"fmt"
)

func main() {
	str := ""

	for _,v := range(os.Args) {
		str += v + " "
	}

	fmt.Println(str)
}

