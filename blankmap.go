package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	a := make(map[string]string)
	a["string"] = "another string"
	a["s"] = "an abbreviation"
	fmt.Println(a)
	a = make(map[string]string)
	fmt.Println(a)
	b := make(map[int]string)
	for i := 0; i <= 10; i++ {
		b[i] = "counting"
		fmt.Println(b)
		if i == 10 {
			b = make(map[int]string)
		}
	}
	fmt.Println(b)

}
