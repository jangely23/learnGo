/* Ejecuta un codigo siempre y cuando el caso coincida con la validación realizada. La instrucción switch comienza con la palabra clave switch y le sigue, en su forma más básica, alguna variable contra la cual puedan realizarse comparaciones. A esto le sigue un par llaves ({}) en el que pueden aparecer varias_ cláusulas de caso_. Las cláusulas de caso describen las acciones que su programa de Go debe realizar cuando la variable proporcionada a la instrucción switch es igual al valor referido por las cláusulas de caso. */

package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}

	for _, valor := range x { //recorre un slice y en cada vuelta valida el caso que sea verdadero

		switch valor {
		case 1:
			fmt.Print("es el uno\n")
		case 2:
			fmt.Print("es el dos\n")
		case 3:
			fmt.Print("es el tres\n")
		case 4:
			fmt.Print("es el cuatro\n")
		case 5:
			fmt.Print("es el cinco\n")
		default:
			fmt.Print("son los numeros")
		}
	}

	switch { //no tiene una expresion a validar por tanto solo ejecutara el primer caso cierto
	case 3 != 5:
		fmt.Print("\nes diferente \n")
	case 4 == 6:
		fmt.Print("\nes igual\n")
	case 4 == 4:
		fmt.Print("\nes menor\n")
	default:
		fmt.Print("son los numeros")
	}
}
