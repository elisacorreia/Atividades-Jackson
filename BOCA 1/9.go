package main

import (
	"fmt"
)

func main() {
	var peso, altura float64
	fmt.Scan(&peso, &altura)

	imc := peso / (altura * altura)

	if imc < 18.5 {
		fmt.Println("Abaixo do peso")
	} else if imc < 25 {
		fmt.Println("Peso normal")
	} else if imc < 30 {
		fmt.Println("Sobrepeso")
	} else {
		fmt.Println("Obesidade")
	}
}
