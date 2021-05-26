package main

import "fmt"

func main() {
	s1 := foo() //ejecuta una funcion que solo retorna un string
	fmt.Println(s1)

	s2 := bar() //ejecuta la funcion de retorno funcion
	fmt.Printf("tipo de dato: %T\n", s2)

	fmt.Println(s2())

}

func foo() string {
	return "Hola mundo" //retorna el string
}

func bar() func() int {

	return func() int { // retorna una funcion
		return 1000
	}
}
