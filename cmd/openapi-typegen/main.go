package main

import (
	"fmt"
	"os"
)

func main() {
	err := NewRootCmd().Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
