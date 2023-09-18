package main

import (
	"fmt"
	"os"
)

func main() {
	for i,v := range os.Environ() {
		fmt.Println(i)
		fmt.Println(v)
		fmt.Println("****")
	}
}
