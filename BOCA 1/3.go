package main

import (
	"fmt"
)

func main() {

	var x, y, z int
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)

	if x+y <= z {
		fmt.Print("Nao forma triangulo")
	} else {
		if x == y && y == z {
			fmt.Println("Equilatero")

		} else if x == y || x == z || y == z {
			fmt.Println("Isosceles")
		} else {
			fmt.Println("Escaleno")
		}
	}
}
