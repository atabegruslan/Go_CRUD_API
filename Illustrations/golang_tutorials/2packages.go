package main

import ("fmt"
		"math"
		"math/rand") // package inside package

func sqrt(x float64) {
	//fmt.Println("SQRT: ", math.Sqrt(x)) // Capitalized functions are the exported functions of the package
	fmt.Printf("SQRT of %f is: %f \n", x, math.Sqrt(x)) // Sprintf wont print to console
	fmt.Printf(`----
----
`)
}

func random(x int) {  // Cant name this function rand
	fmt.Println("Random number: ", rand.Intn(x)) // If you declare function param as eg (x int32), you'll get error: cannot use x (type int32) as type int in argument to rand.Intn
}

// For short-hand: add(x , y float64) , because both x and y have the same type
func add(x float64, y float64) float64 {
	return x+y
}

func main() {
    sqrt(4)
	random(100)
	
	
	//var num1 float64 = 5.1
	//var num2 float64 = 6.2
	
	// Shorter
	//var num1,num2 float64 = 5.1,6.2
	
	// Shorter still
	num1,num2 := 5.1,6.2 // Default float64
	// := is declare and initialize without type. Can only be used inside a function. https://tour.golang.org/basics/10
	
	
	fmt.Println("Sum: ", add(num1,num2))
	
	// pointers
	variable := 5
	address := &variable
	atAddress := *address
	
	fmt.Println("variable: ", variable)
	fmt.Println("address: ", address)
	fmt.Println("atAddress: ", atAddress)
	
	*address = 6
	fmt.Println("variable: ", variable)
	
	*address = *address**address
	fmt.Println("variable: ", variable)
}