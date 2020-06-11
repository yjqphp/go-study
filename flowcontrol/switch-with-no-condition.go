package main

import (
	"fmt"
	"time"
)

/**
	没有条件的switch
	没有条件的switch同switch true一样

	这种形式能将以长串if-then-else写得更加清晰
 */
func main()  {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
