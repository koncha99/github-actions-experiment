package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Running actions...")
	fmt.Println(os.Getenv("github.repository"))
}
