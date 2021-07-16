//canal con bufer no requiere de un receptor

package main

import "fmt"

func main() {
	ca := make(chan int, 1) //el segundo argumento de make es la cantidad de elemntos a tener

	ca <- 63

	fmt.Println(<-ca)
}
