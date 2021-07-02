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

	//crear hash para clave
	bs, err := bcrypt.GenerateFromPassword([]byte(clave), 4) //los argumentos son contraseña en slice de byte y el costo

	if err != nil {
		fmt.Println("error es: ", err)
	}
	fmt.Println(clave)
	fmt.Println(bs)         //muestra el hash en slice de bytes
	fmt.Println(string(bs)) //muesta el hash en caracteres que se almacenaria en una BD

	//verificar si contraseña es correcta

	claveLogin := `clave12345`

	err = bcrypt.CompareHashAndPassword(bs, []byte(claveLogin)) //compara que la has de la claveLogin sea igual al de bs y retorna err en caso de ser diferente

	if err != nil {
		fmt.Println("clave errada")
		return
	} else {
		fmt.Println("login exitoso")
	}

}
