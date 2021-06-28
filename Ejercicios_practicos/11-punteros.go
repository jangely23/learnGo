package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x befor", &x) //puntero
	fmt.Println("x befor", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y) //de referencia *y para acceder a el valor de esa direccion de memoria
	*y = 43                    //modifica el valor contenido en la direccion de memoria de y
	fmt.Println("y after", y)
}
