package main

import (
    "fmt"
    "reflect"
)

func main() {
    var a float64 = 2
    fmt.Println(a/2)
    fmt.Println(reflect.TypeOf(a/2))
}