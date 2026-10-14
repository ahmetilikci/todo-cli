package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Access denied")
		return
	}

	command := os.Args[1]
	switch command {
	case "add":
		fmt.Println("Added", os.Args[2])
	case "list":
		fmt.Println("There's not any task")
	default:
		fmt.Println("Unknown command", command)

	}
}
