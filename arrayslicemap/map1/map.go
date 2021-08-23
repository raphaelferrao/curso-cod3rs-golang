package main

import "fmt"

func main() {
	// var aprovados map[int]string -> nao funciona, mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[223456789] = "Pedro"
	aprovados[323456789] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}

	fmt.Println(aprovados[123456789])
	delete(aprovados, 123456789)
	fmt.Println(aprovados[123456789])
	fmt.Println(aprovados)
}
