/* son valores que una vez definidos no pueden cambiar, en Go existen con tipo y sin tipo, para crearlos se usa la palabla clave const.  */

package main

import "fmt"

const (
	year     = 365        //sin tipo
	leapYear = int32(366) //con tipo
)

func main() {
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)
	fmt.Println(minutes * year)
	fmt.Println(minutes * leapYear)
}
