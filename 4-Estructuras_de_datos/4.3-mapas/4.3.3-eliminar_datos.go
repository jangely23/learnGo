package main

import (
	"fmt"
)

func main() {

	x := map[string][]string{

		`jessica`: []string{`arte`, `cine`, `musica`},
		`danelly`: []string{`deporte`, `cine`, `musica`},
		`lauren`:  []string{`arte`, `cine`, `musica`},
	}

	x[`martha`] = []string{`cocinar`, `pintar`, `tejer`}

	delete(x, "jessica") //borra un registro de un mapa

	for llave, valor := range x {
		fmt.Printf("\nPara %v, sus cosas favoritas son: %v\n", llave, valor)

		for indice, contenido := range valor {
			fmt.Printf("\t%v - %v\n", indice+1, contenido)
		}
	}
}
