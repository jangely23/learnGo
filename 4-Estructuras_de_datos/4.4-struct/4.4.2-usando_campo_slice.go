package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	helados  []string //slice como campo
}

func main() {
	p := persona{
		nombre:   "jessica",
		apellido: "leonel",
		helados:  []string{`chocolate`, `browine`, `mora`},
	}

	fmt.Println(p.nombre, p.apellido)
	for indice, valor := range p.helados {
		fmt.Println("\t", indice, valor)
	}
}
