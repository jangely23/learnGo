//canal con bufer con fallo

package main

import "fmt"

func main() {
	ca := make(chan int, 1) //el segundo argumento de make es la cantidad de elemntos a tener

	ca <- 63
	ca <- 42 //estoy enviando mas de lo que el bufer almacena

	fmt.Println(<-ca)
}
