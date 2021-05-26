package main

import (
	"fmt"
)

func main() {

	x := make([]string, 5, 5) //tipo slice, con logitud de 2 y capacidad de 5

	x = []string{`jessica`, `lauren`, `martha`, `mariana`, `zully`}

	fmt.Println(len(x)) //funcion que imprime longitud
	fmt.Println(cap(x)) //funcion que imprime capacidad

	for i := 0; i < len(x); i++ {
		println(i, x[i])
	}

}
