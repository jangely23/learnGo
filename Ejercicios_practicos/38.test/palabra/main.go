//Package palabra nos ayuda con opreaciones de srings
package palabra

import "strings"

//COnteoUso retorna un mapa con las palaras y cantidad de veces qu parece
// No necesitas escribir un ejemplo para esta función
// escribir una prueba para esta es un bonocdcd
 retador; un poco más difícil
func ConteoUso(s string) map[string]int {
	xs := strings.Fields(s) //obtiene cada plabra y los pone en el slice xs
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

func Conteo(s string) int {
	// escribe el código para esta función
	xs := strings.Fields(s) //obtiene cada plabra y los pone en el slice xs

	return len(xs)
}
