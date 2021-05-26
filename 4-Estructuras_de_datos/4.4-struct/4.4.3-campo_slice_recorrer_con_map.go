package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	helados  []string //slice como campo
}

func main() {

	p1 := persona{
		nombre:   "Jessica",
		apellido: "Leonel",
		helados:  []string{`Frutos rojos`, `Browni`, `Chicle`},
	}

	p2 := persona{
		nombre:   "Lauren",
		apellido: "Laguna",
		helados:  []string{`Frutos rojos`, `Browni`, `Chicle`},
	}

	m := map[string]persona{ //crea un mapa con llaves string y valores tipo persona
		p1.apellido: p1,
		p2.apellido: p2,
	}

	for _, v := range m { //guion bajo desecha una variable que se requiere pero no se usara
		fmt.Println(v.nombre, v.apellido)

		for indice, valor := range v.helados {
			fmt.Println(indice, ").", valor)
		}
	}
}
