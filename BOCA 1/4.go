package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	if i >= 5 && i <= 7 {
		fmt.Println("Infantil A")
	} else if i >= 8 && i <= 10 {
		fmt.Println("Infantil B")
	} else if i >= 11 && i <= 13 {
		fmt.Println("Juvenil A")
	} else if i >= 14 && i <= 17 {
		fmt.Println("Juvenil B")
	} else if i >= 18 {
		fmt.Println("Adulto")
	}

}
