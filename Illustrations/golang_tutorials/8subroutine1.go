package main

import ("fmt"
		"time")

func say(s string) {
	for i := 1; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
}

func main() {
	go say("AAAAAA") // Parallel sub routine. If program finishes before this, then the program will stop without waiting for this
	say("BBBBBB")
}
