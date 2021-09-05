package main

import "fmt"

type car struct {
	gas uint16
	brake uint16
}

func (c car) gasMinusBrake() float64 { // Value Receiver - pass by value
	return float64(c.gas) - float64(c.brake)
}

func (c *car) newGas(gasNew float64) { // Pointer Receiver - pass by ref
	c.gas = uint16(gasNew)
}

func main() {
	car1 := car{gas:1,brake:2}
    fmt.Println(car1.gas)
    fmt.Println(car1.gasMinusBrake())
	car1.newGas(3)
    fmt.Println(car1.gas)
    fmt.Println(car1.gasMinusBrake())
}