/* zona de almacenamiento contiguo que contiene una serie de elementos del mismo tipo, los elementos de la matriz. ​ Desde el punto de vista lógico una matriz se puede ver como un conjunto de elementos ordenados en fila (o filas y columnas si tuviera dos dimensiones) */

package main

import (
	"fmt"
)

func main() {
	ct := [5]int{10, 12, 12, 435, 67}

	for clave, valor := range ct {
		fmt.Printf("Tipo de arreglo: %T, posicion %d, contenido %d\n", ct, clave, valor)
	}

	mv := [2]string{"jessica", "lauren"}

	for indice, valor := range mv {
		fmt.Printf("Tipo de arreglo: %T, posicion %d, contenido %s\n", mv, clave, valor)
	}
}
