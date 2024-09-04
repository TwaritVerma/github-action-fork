package main

import (
    "fmt"
	"os"
)

func main() {
	secret := os.Getenv("SECRET_KEY")
	fmt.Println(secret)
	if secret != "SECRET_VALUE" {
		fmt.Println("WRONG SECRET")
		os.Exit(1)
	}
	fmt.Println("Completed")
}
