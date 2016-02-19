// Working off code from https://blog.golang.org/pipelines
package main

import (
	"fmt"
	"reflect"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func main() {
	c := gen(2, 3)
	out := sq(c)

	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(out))
	fmt.Println(<-out)
	fmt.Println(<-out)
}
