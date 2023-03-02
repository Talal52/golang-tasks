package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	talal, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	strings.Fields(string(talal))
	fmt.Printf("length = %d",len(talal))
}
