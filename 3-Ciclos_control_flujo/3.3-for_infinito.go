/* Cuando se usa este ciclo no se evalua un condicion para salir del ciclo, por el contrario requiere de  la palabra clave break para salir de el. */

package main

import "fmt"

func main() {
	i := 0

	for {
		if i > 3 {
			break
		}

		fmt.Println("El valor es:", i)

		i++
	}
}
