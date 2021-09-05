package main

import (
    "fmt"
    "sync"
    "time"
)

// Full example https://gobyexample.com/waitgroups

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil { // Like a catch block
		fmt.Println("Recovered in cleanup", r)
	}
}

func say(s string) {
	// Deferred code waits for everything in its function to finish or error first
	defer wg.Done() // 1) evaluates Done func. 5) Executes Done
	defer cleanup() // 2) evaluates cleanup func. 4) Executes cleanup
	for i := 1; i < 5; i++ { // 3) Executes the loop
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
		if i == 2 {
			panic("Example error") // Like throw Exception("...")
		}
	}
}

func main() {
	wg.Add(1)
	go say("AAAAAA")
	wg.Add(1)
	go say("BBBBBB")
	wg.Wait()
}
