/* Las variables se pueden crear para ser usadas de forma global entro del programa o local dentro del scope de una función, para usarlas de forma global se crean y asignan antes de la **funcion main** usando la palabra clave **var,** el operador de asignación **=** para inicializar la variable, el valor a asignar y el tipo de dato, por **ejemplo:** **var a int = 5.**

Para asignar la variable dentro del scope de la función se debe inicializar la primera vez omitiendo la palabra var y usando el operador de declaración corto "**:=**", este operador solo se usa la primera vez, posteriormente solo se usara el operador de asignación "=" para cambiar el valor de la variable, ejemplo de asignación: **e = 42**. */

package main

import (
	"fmt"
)

var a string = "Hola," //creacion y asignacion global
var c bool

func main() {
	b := "mundo." //creacion y asignacion local scope funcion main
	fmt.Println(a, b)
	fmt.Println(c)
}
