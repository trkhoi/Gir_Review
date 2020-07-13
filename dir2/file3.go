package main

import (
	"fmt"
	"os"
)

func main() {
	argName := os.Args[1]
	fmt.Printf("the argument is %s", argName)
}
