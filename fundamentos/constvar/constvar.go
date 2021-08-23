package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64 inferido pelo compilador

	// forma reduzida de criar uma var com valor inicial
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a + b + c + d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa"
	fmt.Println(g, h, i)

	type Pessoa struct {
		Nome  string
		Age   uint8
		Email string
	}

	p1 := Pessoa{"Strawlley", 31, "straw@email.com"}
	p2 := Pessoa{"Strawlley", 31, "straw@email.com"}

	fmt.Println(p1 == p2)
}
