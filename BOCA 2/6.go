package main

import "fmt"

func main(){
	var n, i int
	var p1, p2 float64
	
	
	fmt.Scan(&n)

	
	i = 0
	media := 0
	reprov := []int{}
	exame := []int{}
	aprov := []int{}
	
	for i <= n {
		i++
		fmt.Scan(&p1)
		fmt.Scan(&p2)
		media = (p1 + p2) / 2
		if media <= 3{
			fmt.Println("Aluno ", i, ": Reprovado")
			reprov = append(reprov, i)

		} else if media > 3 && media < 7 {
			fmt.Println("Aluno ", i, ": Exame")
			exame = append(exame, i)
		} else {
			fmt.Println("Aluno ", i, ": Aprovado")
			aprov = append(aprov, i)
		}
	} // fazer um range com os indices do slice
}
