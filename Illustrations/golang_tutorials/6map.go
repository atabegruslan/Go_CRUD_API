package main

import "fmt"

// https://blog.golang.org/maps

func main() {
	grades := make(map[string]int) // The make function allocates and initializes a hash map data structure and returns a map value that points to it.
	
	grades["tim"] = 42 // Cant use single quotes
	
    fmt.Println("Tim: ", grades["tim"])
}