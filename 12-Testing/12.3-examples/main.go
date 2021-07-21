package main

import (
	"fmt"

	"github.com/learnGo/12-Testing/12.3-examples/mate"
)

func main() {
	fmt.Println(mate.Sum(2, 4, 5))
	fmt.Println(mate.Sum(4, 7, 8, 0))
}
