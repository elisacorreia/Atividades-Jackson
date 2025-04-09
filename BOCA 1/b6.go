package main

import ("fmt")

func main() {
	var conta int
	var lim, ini, dep, ret, resul float64

	fmt.Scan(&conta)
	fmt.Scan(&lim)
	fmt.Scan(&ini)
	fmt.Scan(&dep)
	fmt.Scan(&ret)

	resul = (ini + dep) - ret

	if resul > lim {
		fmt.Println("Credito excedido: ", resul)
	} else {
		fmt.Println("Saldo: ", resul)
	}
}