package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 1; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
	wg.Done() // https://golang.org/pkg/sync/#WaitGroup.Done counter--
}

func main() {
	wg.Add(1) // https://golang.org/pkg/sync/#WaitGroup.Add counter++
	go say("AAAAAA")
	wg.Add(1)
	go say("BBBBBB")
	wg.Wait() // https://golang.org/pkg/sync/#WaitGroup.Wait waits for counter to be 0
}
