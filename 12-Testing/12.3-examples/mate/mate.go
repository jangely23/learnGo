//Package mate comprueba si sabes sumar
package mate

//Sum suma una cantidad ilimitada de numeros enteros
func Sum(xi ...int) int {
	var s int

	for _, v := range xi {
		s += v
	}

	return s
}
