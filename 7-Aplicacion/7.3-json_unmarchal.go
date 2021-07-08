/*
Convierte un Json a un tipo de dato entendido por el programa, ejemplo Struct
se usara la funcion unmarchaling que recibe un Json tipo byte y un tipo de dato cualquiera, retorna error.
el unmarcheling lo hace a un puntero del valor ejemplo &bs
*/

package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	// uso de alias json
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	Edad     int    `json:"Edad"`
}

func main() {
	//json
	s := `[{"Nombre":"jessica","Apellido":"leonel","Edad":31},{"Nombre":"lauren","Apellido":"laguna","Edad":14}]`
	//conversion
	bs := []byte(s) // creamos un slice de byte que contiene a el literal s

	//imprime los tipos
	fmt.Printf("la s es %T\n", s)
	fmt.Printf("bs es %T\n", bs)

	//personas := []persona{} otra forma de escribir esto mismo y mas legible es la siguiente:
	var personas []persona

	err := json.Unmarshal(bs, &personas) //realiza unmarchal

	//manejando errores
	if err != nil {
		fmt.Println("error es: ", err)
	}

	fmt.Println("\nToda la data", personas)

	//recorriendo el slice personas
	for i, v := range personas {
		fmt.Println("\nPersona numero: ", i+1)
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
	}
}
