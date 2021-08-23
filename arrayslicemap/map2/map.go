package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Joao":      11325.40,
		"Gabriela Silva": 12445.12,
		"Pedro Junior":   1200.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Joao Tortugo") //nao causa erro deletar alguma chave que nao existe

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	for nome := range funcsESalarios {
		fmt.Println(nome)
	}
}
