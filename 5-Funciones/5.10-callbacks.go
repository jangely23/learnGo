/* son funciones que se pasan como parametro a otra funcion, se usa en programacion funcional y no es recomendado en Go*/

package main

import "fmt"

func main() {
	enteros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	resultado1 := suma(enteros...) //los tres puntos envian los valores individuales como argumento
	fmt.Println(resultado1)

	resultadoPares := sumaPar(suma, enteros...) // hace callback a sumaPar
	fmt.Println(resultadoPares)

	resultadoImpares := sumaImpar(suma, enteros...) // hace callback a sumaImpar
	fmt.Println(resultadoImpares)

}

func suma(xi ...int) int { // recibe como parametro los valores individuales de un slice
	total := 0

	for _, valor := range xi { // realiza la suma del slice
		total += valor
	}

	return total
}

func sumaPar(f func(xi ...int) int, vi ...int) int { //recibe como parametro una funcion y el slice
	var y []int
	for _, valor := range vi {
		if valor%2 == 0 {
			y = append(y, valor)
		}
	}
	return f(y...)
}

func sumaImpar(f func(xi ...int) int, vi ...int) int { //recibe como parametro una funcion y el slice
	var y []int
	for _, valor := range vi {
		if valor%2 != 0 {
			y = append(y, valor)
		}
	}
	return f(y...)
}
