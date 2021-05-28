// crear un closure, una funcion foo que retorna un entero de funcion anonima

package main

import "fmt"

func main() {
	f := foo()
	j := foo()

	//copia el entorno de f, mantiene o persiste el valor de f
	fmt.Println("f vale: ", f())
	fmt.Println("f vale: ", f())
	fmt.Println("f vale: ", f())
	fmt.Println("f vale: ", f())
	fmt.Println("f vale: ", f())

	//copia el entorno de j, mantiene o persiste el valor de j
	fmt.Println("j vale: ", j())
	fmt.Println("j vale: ", j())
	fmt.Println("j vale: ", j())
	fmt.Println("j vale: ", j())
	fmt.Println("j vale: ", j())

}
func foo() func() int {
	x := 100
	return func() int {
		x++
		return x
	}
}
