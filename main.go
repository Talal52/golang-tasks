package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(files)
	for i := 0; i < 2; i++ {
		fmt.Println(files[i].Name())
	}
}
