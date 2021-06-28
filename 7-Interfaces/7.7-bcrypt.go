/*
Bcrypt es una herramienta de proteccion de datos, ofrece encriptacion.
funcion de has: Aplica un algoritmo matematico para encriptar un texto que
se pasa como algumento a una funcion de hasching

Regularmente se utiliza para guardar la contraseña de los usuarios.
El costo de la funcion tiene como parametro un costo (poder de computo para romper el password).

Es un paquete externo al core de Go
*/

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	clave := `clave123`

	bs, err := bcrypt.GenerateFromPassword([]byte(clave), 4) //los argumentos son contraseña en slice de byte y el costo

	if err != nil {
		fmt.Println("error es: ", err)
	}
	fmt.Println(clave)
	fmt.Println(bs)

}
