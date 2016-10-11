package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	cmd := os.Args[1]
	fmt.Printf("Given date: %v\n", cmd)
	t, _ := time.Parse("2006-01-02", cmd)
	fmt.Printf("Unix nano time: %v\n", t.UnixNano())
	fmt.Printf("Unix time: %v\n", t.Unix())
}
