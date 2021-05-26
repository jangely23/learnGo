/* Los slices pueden verse como arreglos de longitud dinámica, siendo un poco más técnicos, un slice apunta a un array */

package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 45, 7, 6, 3, 4, 5, 7}

	for indice, valor := range x {
		fmt.Printf("Slice de tipo: %T\tcon indice %d\t de valor %d\n", x, indice, valor)
	}
}
