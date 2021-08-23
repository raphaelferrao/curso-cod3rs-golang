package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(reflect.TypeOf(a1), a1)
	fmt.Println(reflect.TypeOf(s1), s1)

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice nao eh um array! Slice define um pedaco de um array
	s2 := a2[1:3]
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(a2, s3)
}
