package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereco da variavel i

	*p++ // desreferenciando (pegando o valor)
	fmt.Println(p, *p, i, &i)

	i++
	fmt.Println(p, *p, i, &i)

	// Go nao tem aritmetica de ponteiros
	// p++
}
