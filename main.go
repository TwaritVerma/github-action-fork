package main

import (
    "fmt"
	"os"
)

func main() {
	secret := os.Getenv("FORK_SECRET")
	fmt.Println(secret)
	if secret != "mySecret" {
		fmt.Println("WRONG SECRET")
		os.Exit(1)
	}
	fmt.Println("Completed")
}
