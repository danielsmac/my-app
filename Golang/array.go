package main

import "fmt"

func main()  {
	
	// criar array com 5 ints. o tipo do elemento e o tamanho sao partes do tipo do array.

	var a [5]int
	fmt.Println("emp:", a)

	// podemos usar um indice com valor para o array[index]
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// o builtin len retorna o tamanho do array
	fmt.Println("len: ", len(a))

	// podemos usar a sintax para declarar e inicializar um array com uma unica linha
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// pode-se também compilar um contador de elementos usando [...]
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//se especificar com (:), os elementos entre serão 0
	b = [...]int{100, 3:400, 500}
	fmt.Println("idx:", b)

	//Array não unidimensionais, mas podem compor tipos de estruturas de dados multi-dimensionais
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int {
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}