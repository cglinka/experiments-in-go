package main

import (
	"fmt"
	"time"
)

func main() {
	// My goal here is to gain a better understanding of how the time.Time package works.
	// I want to compare various timestamps in different foramts, so I need to understand
	// what I'm doing in order to make valid comparisons.
	t := time.Now()
	fmt.Println("now: ", t)

	// This doens't work b/c mismatched types (time.Time and time.Duration):  t30 := t - time.Second*30
	// This doens't work b/c mismatched types (int64 and time.Duration):  t30 := t.Unix() - time.Second*30

	tUnix := t.Unix()
	fmt.Println("Unix time: ", tUnix)

	// This doesnt' work because  mismatched types (time.Time and time.Duration): t30 := tUnix - 30*time.Second

	t30 := tUnix - 30
	fmt.Println("30 seconds ago: ", t30)

	t50 := tUnix - 50
	fmt.Println("50 seconds ago: ", t50)

	// This doesn't work because t50 is type int64, not a time type: diff := t50.Sub(tUnix)
	// Same for tUnix.Sub(t50)
	// This works because we are calling Sub on a time.Time and converting t50 to a time.Time from int64
	diff := t.Sub(time.Unix(t50, 0))
	fmt.Println("t50 sub tunix: ", diff)

	// Lets do this for real now.
	// Get current time:
	tnow := time.Now()
	fmt.Println("tnow: ", tnow)
	tearlier := tnow.Add(-31 * time.Second)
	fmt.Println("tearlier: ", tearlier)
	//  This forkulation results in a negative duration: diff = tearlier.Sub(tnow)
	diff = tnow.Sub(tearlier)
	fmt.Println("diff: ", diff)
	if diff <= 30*time.Second {
		fmt.Println("Less than or equal to 30 seconds")
	} else {
		fmt.Println("greater than 30 seconds")
	}

}
