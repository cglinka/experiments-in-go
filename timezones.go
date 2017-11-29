package main

import (
	"fmt"
	"time"
)

// How do timeszones work anyway?
// https://www.goinggo.net/2013/08/using-time-timezones-and-location-in-go.html

func main() {
	loc, err := time.LoadLocation("America/Chicago")
	fmt.Println(loc)
	fmt.Println(err)
	loc, err = time.LoadLocation("America/Dallas")
	fmt.Println(loc)
	fmt.Println(err)

	loc, err = time.LoadLocation("America/Houston")
	fmt.Println(loc)
	fmt.Println(err)
	loc, err = time.LoadLocation("America/StLouis")
	fmt.Println(loc)
	fmt.Println(err)
}
