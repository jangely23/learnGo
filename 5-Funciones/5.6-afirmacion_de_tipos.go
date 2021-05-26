package main

import (
	"fmt"
)

type figura interface { //crea una interfaz

	area() int //metodos que usa la interfaz

}

type cuadrado struct {
	lado int
}

type triangulo struct {
	base   int
	altura int
}

func main() {

	c := cuadrado{ // crea un cuadrado
		lado: 50,
	}

	t := triangulo{ // crea un triangulo
		base:   10,
		altura: 20,
	}

	//usamos la interfaz
	verArea(c)
	verArea(t)

}

func (ac cuadrado) area() int { //metodo area de cuadrado
	area := ac.lado * ac.lado
	return area
}

func (at triangulo) area() int { //metodo area de un triangulo
	area := (at.base * at.altura) / 2
	return area
}

func verArea(f figura) { //pasamos la interfaz como parametro en la funcion

	//con afirmacion de tipos
	switch f.(type) { //Permite la afirmacion de tipos para aceder a los campos de cada struct

	case cuadrado: // en caso de que sea de este tipo

		//utilizamos el metodo de la interfaz y los campos del struct
		fmt.Printf("El area del cuadrado con lado %vcm es: %vcm\n", f.(cuadrado).lado, f.area())

	case triangulo:

		fmt.Printf("El area del triangulo con base %vcm y altura %vcm es: %vcm\n", f.(triangulo).base, f.(triangulo).altura, f.area())
	}

	//sin afirmacion de tipos
	fmt.Printf("el area es: %v\n\n", f.area())
}
