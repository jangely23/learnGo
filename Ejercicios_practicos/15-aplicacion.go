//hacer unmarchal a un json hacia una structura de datos de GO

package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre   string   `json:"Nombre"`
	Apellido string   `json:"Apellido"`
	Edad     int      `json:"Edad"`
	Dichos   []string `json:"Dichos"`
}

func main() {
	s := `[{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Nombre":"M","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`
	fmt.Println("\n", s)

	var personas []Persona                      //crea la variable interfaz personas de tipo slice de Persona
	err := json.Unmarshal([]byte(s), &personas) //se pasa un puntero a la interfaz
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println("\n", personas, "\n")

	for _, v := range personas {
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
		for _, c := range v.Dichos {
			fmt.Println("\t", c)
		}
		fmt.Println("")
	}
}
