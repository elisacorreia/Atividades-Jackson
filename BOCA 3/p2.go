package main

import "fmt"

func main(){
	var n int

	fmt.Scan(&n)
	if n >= 9 {
		fmt.Println("A")
	} else if n >=7 && n < 9{
		fmt.Println("B")
	} else if n >=5 && n < 7{
		fmt.Println("C")
	} else if n >=3 && n < 5{
		fmt.Println("D")
	} else if n < 3 {
		fmt.Println("E")
	}
}