/* Un map nos permite almacenar información de forma similar a como lo hacíamos con los arreglos o los slices pero con la particularidad de que cada elemento de un map tiene asociada una llave única */

package main

import (
	"fmt"
)

func main() {

	x := map[string][]string{ //dentro de los corchetes tipo de la llave, fuera tipo de los valores
		`jessica`: []string{`arte`, `cine`, `musica`},
		`danelly`: []string{`deporte`, `cine`, `musica`},
		`lauren`:  []string{`arte`, `cine`, `musica`}, //importante la coma entre elementos del mapa
	}

	for llave, valor := range x {
		fmt.Printf("\nPara %v, sus cosas favoritas son: %v\n", llave, valor)

		for indice, contenido := range valor {
			fmt.Printf("\t%v - %v\n", indice+1, contenido)
		}
	}
}
