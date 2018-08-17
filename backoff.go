package main

import (
	"fmt"
	"time"

	"github.com/jpillora/backoff"
)

func main() {
	b := &backoff.Backoff{
		Min:    time.Second,
		Max:    30 * time.Second,
		Jitter: true,
	}
	fmt.Println("Going to start for loop...")

	// ForLoop:
	// for {
	// 	fmt.Println("PRINT FORLOOP")
	// 	duration := b.Duration()
	// 	fmt.Println("duration: ", duration)
	// 	select {
	// 	case <-time.After(duration):
	// 		break ForLoop
	// 	}
	// }

	for {
		fmt.Println("PRINT")
		duration := b.Duration()
		fmt.Println("duration: ", duration)
		select {
		case <-time.After(duration):
			continue
			// case <-time.After(duration):
			// 	return
		}
	}

	for {
		fmt.Println("last loop")

		if true {
			break
		}
		fmt.Println("Did we break?")
	}
}
