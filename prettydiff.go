package main

import (
	"fmt"

	"github.com/kr/pretty"
)

func main() {

	a := map[string]string{"a": "a"}
	b := map[string]string{"a": "b"}
	diff := pretty.Diff(a, b)
	fmt.Println(diff)
}
