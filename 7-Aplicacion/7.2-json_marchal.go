/*Marchal de datos en Go a JSON
Importante los cambos del struc deben iniciar en Mayuscula.
Marshal devuelve la codificación JSON
Importante los campos debe iniciar con mayuscula para que se puedan marchear
*/

package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string //campos en mayuscula la primera letra
	Apellido string
	Edad     int
}

func main() {
	p1 := persona{
		Nombre:   "jessica",
		Apellido: "leonel",
		Edad:     31,
	}

	p2 := persona{
		Nombre:   "lauren",
		Apellido: "laguna",
		Edad:     14,
	}

	//creamos una variable de tipo slice persona que almacene los valores p1 y p2
	personas := []persona{p1, p2}

	fmt.Println(personas)

	elaborado, err := json.Marshal(personas) //funcion convierte a Json Un dato y devulve 2 resultados la conversion (bytes) y el error.

	if err != nil { //control de errores

		fmt.Println("error es: ", err)
	}

	fmt.Println(string(elaborado)) //muestra el bytes convertido a string para que sea entendible.
}
