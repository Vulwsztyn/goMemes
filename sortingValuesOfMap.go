package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make(map[int] []int)
	a[1] = []int{3,2,1}
	sort.Ints(a[1])
	fmt.Println(a[1])
}
