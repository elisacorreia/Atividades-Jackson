package main

import "fmt"

func main() {
	var vendas float64
	fmt.Scan(&vendas)

	salarioBase := 500.0
	comissao := vendas * 0.09
	bonus := 0.0

	if vendas > 15000 {
		bonus = 800.0
	}

	salarioFinal := salarioBase + comissao + bonus

	fmt.Printf("%.5f\n", salarioFinal)
}
