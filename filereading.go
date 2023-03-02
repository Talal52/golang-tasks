package main

import (
	"fmt"
	"os"
)

func main() {
	talal, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(talal))
}
