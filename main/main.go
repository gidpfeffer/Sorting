package main

import (
	"fmt"
	"github.com/gidpfeffer/sorting/quicksort"
)

func main() {
	s := []int{110, 2, 23, 45, 23, 12, 32, 12, 43, 54, 56, 67, 34, 56, 67}
	quicksort.Sort(s)
	fmt.Println(s)
}
