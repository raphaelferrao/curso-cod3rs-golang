package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	// array1 = append(array1, 4, 5, 6) nao eh permitido
	slice1 = append(slice1, 4, 5, 6) //expande o slice original e insere os novos elementos
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // so copia os itens do tamanho do slice destino
	fmt.Println(array1, slice1, slice2)

}
