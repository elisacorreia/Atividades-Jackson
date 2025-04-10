package main

import "fmt"

func main() {
	var ano int
	fmt.Scan(&ano)

	if (ano%4 == 0 && ano%100 != 0) || (ano%400 == 0) {
		fmt.Println("Bissexto")
	} else {
		fmt.Println("Nao bissexto")
	}
}
