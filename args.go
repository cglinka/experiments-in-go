package main

import (
	"fmt"
	"os"
)

func main() {
	in := os.Args[0]
	in1 := os.Args[1]
	fmt.Println("Args[0]: ", in)
	fmt.Println("Args[1]: ", in1)
}
