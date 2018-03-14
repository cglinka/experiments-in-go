package main

import (
	"fmt"

	"github.com/kr/pretty"
)

func main() {

	a := map[string]string{"a": "a"}
	b := map[string]string{"a": "a"}
	diff := pretty.Diff(a, b)
	fmt.Println(diff)
	fmt.Println(len(diff))
}
