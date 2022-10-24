package main

import (
	"fmt"
	"os"
)

func main() {
	hasPass := false
	for i := range os.Args {
		if os.Args[i] == "password" {
			hasPass = true
			break
		}
	}
	if !hasPass {
		panic("shit went wrong!")
	}
	fmt.Println("success")

}
