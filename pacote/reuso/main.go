package main

import (
	"fmt"

	"github.com/raphaelferrao/area"
)

// Para dar certo foi necessário rodar
// go env -w GO111MODULE=auto
func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
}
