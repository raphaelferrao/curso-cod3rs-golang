package main

import "fmt"

func inc1(n int) {
	n++
}

// um ponteiro eh representado por um *
func inc2(n *int) {
	// * eh usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)
	fmt.Println(n)

	// & usado para obter o endereco da variavel
	inc2(&n) // por referencia
	fmt.Println(n)
}
