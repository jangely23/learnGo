// asignar una funcion a una variable y luego llamar a una funcion.

package main

import "fmt"

func main() {

	numeros := func() {
		for i := 0; i <= 50; i++ {
			fmt.Println(i)
		}
	}
	numeros()

	fmt.Println("finalizado")
}
