package decir

import "fmt"

//Saludar permite saludar a una persona
func Saludar(n string) string {
	return fmt.Sprint("Bienvenido querido ", n)
}
