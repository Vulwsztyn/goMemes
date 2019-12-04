package main

import (
    "fmt"
)
type Interval = [2]float64
func main() {
    var a int
    var b float64
    var c Interval
    if a == 0 {
	fmt.Println("test int 0")
    } else {
	fmt.Println("test int 0n't")
    }
    if b == 0 {
	fmt.Println("test float64 0")
    } else {
	fmt.Println("test float64 0n't")
    }
    if (c == Interval{}) {
	fmt.Println("test Interval 0")
    } else {
	fmt.Println("test Interval 0n't")
    }
}
