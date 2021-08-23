package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 1234.0,
			"Gustavo Silva":  1235.0,
		},
		"J": {
			"Joao Silva":   2234.0,
			"Joel Tortugo": 2235.0,
		},
		"S": {
			"Strawlley Pessoa": 4234.0,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "G")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
