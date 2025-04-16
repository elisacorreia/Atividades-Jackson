package main

import "fmt"

func main() {
	var n, i int

	for {
		fmt.Scan(&n) 
		if n <= 0 { 
			break
		}

		quad := false
		i = 1

		for i*i <= n {
			if i*i == n {
				quad = true
				break
			}
			i++
		}

		if quad {
			fmt.Printf("%d eh quadrado perfeito\n", n)
		} else {
			fmt.Printf("%d nao eh quadrado perfeito\n", n)
		}
	}
}