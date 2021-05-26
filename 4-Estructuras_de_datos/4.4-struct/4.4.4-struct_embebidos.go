package main

import (
	"fmt"
)

//crea un tipo de dato compuesto struct

type persona struct {
	nombre string
	edad   int
	genero string
}

type empleado struct {
	persona //tiene embebido el struct persona
	activo  bool
}

func main() {

	e1 := empleado{
		persona: persona{
			nombre: "jessica", //campos del struct embebido pasan a ser parte del struct externo
			edad:   31,
			genero: "femenino",
		},

		activo: true,
	}

	fmt.Println(e1)
	fmt.Println(e1.nombre, e1.edad, e1.genero, e1.activo) // se acceden directamente los campos del struct persona
}
