package main

import "fmt"

func main() {
	fmt.Println("la suma de 2 + 4 es:  ", suma(2, 4))
	fmt.Println("la suma de 1 + 5 es:  ", suma(1, 5))
	fmt.Println("la suma de 6 + 7 es:  ", suma(6, 7))
}

func suma(n ...int) int {
	var sum int
	for _, v := range n {
		sum += v
	}
	return sum
}
