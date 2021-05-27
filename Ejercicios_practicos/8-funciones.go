// crear una funcion que retorne una funcion

package main

import "fmt"

func main() {
	prueba := foo(10)
	fmt.Println(prueba())
}

func foo(n float64) func() float64 {
	return func() float64 {
		return n * 2
	}
}
