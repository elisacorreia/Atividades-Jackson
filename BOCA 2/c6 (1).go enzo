package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	arr1 := make([]float64, N)

	for i := 0; i < N; i++ {
		var n1, n2 float64
		fmt.Scan(&n1, &n2)
		med := (n1 + n2) / 2
		arr1[i] = med

	}

	totr := 0
	tota := 0
	tote := 0

	for i := 0; i < N; i++ {
		x := i + 1
		if arr1[i] <= 3 {

			fmt.Printf("Aluno %d: Reprovado\n", x)
			totr++
		} else if arr1[i] > 3 && arr1[i] < 7 {
			fmt.Printf("Aluno %d: Exame\n", x)
			tote++
		} else {
			fmt.Printf("Aluno %d: Aprovado\n", x)
			tota++
		}
	}
	fmt.Println("Total Aprovados:", tota)
	fmt.Println("Total Exame:", tote)
	fmt.Println("Total Reprovados:", totr)
}
