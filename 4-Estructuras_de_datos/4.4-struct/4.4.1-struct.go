/* Struct en Golang es una recopilación de datos. Podemos usarlo para definir tipos de datos personalizados. Una estructura puede contener campos miembros dentro de ellos e incluso puede contener otros tipos de estructura. */

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

func main() {
	p1 := persona{ //crea una variabe que contiene un struct
		nombre: "jessica",
		edad:   31,
		genero: "femenino",
	}

	p2 := persona{
		nombre: "johan",
		edad:   1,
		genero: "masculino",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.nombre, p2.edad, p2.genero) // tambien se puede acceder a un campo especifico
}
