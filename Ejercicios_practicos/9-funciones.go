/*crear una funcion foo que se le van a pasar 2 parametros (otra funcion y un slice de enteros)
va a llamar la funcion parametro y al resultado le sumara 1
la funcion parametro obtendra la longitud del slice
si es 0 la longitud entonces retornara 0
si es 1 entonces retorne el valor del slice en el indice 0
si es diferente retornalra la suma del indice 0 y el indice longitud menos 1*/

package main

import "fmt"

func main() {
	x := []int{100, 200, 300, 400}
	resultado := foo(proceso, x)

	fmt.Println(resultado)

}

func foo(f func(s2 []int) int, s []int) int {
	retorno := f(s)
	retorno++
	return retorno
}

func proceso(s2 []int) int {
	if len(s2) == 0 {
		return 0
	} else if len(s2) == 1 {
		return s2[0]
	} else {
		return s2[0] + s2[len(s2)-1]
	}

}
